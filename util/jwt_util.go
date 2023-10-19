package util

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	apiCommon "github.com/loongkirin/gpaas/api/core"
	app "github.com/loongkirin/gpaas/app"
)

type JWTUtil struct {
	SecretKey   []byte
	ExpiresTime int64
	BufferTime  int64
	Issuer      string
}

var (
	TokenExpired   = errors.New("Token is expired")
	TokenNotValid  = errors.New("Token not active yet")
	TokenMalformed = errors.New("That's not even a token")
	TokenInvalid   = errors.New("Couldn't handle this token:")
)

func NewJWTUtil() *JWTUtil {
	return &JWTUtil{
		SecretKey:   []byte(app.AppContext.APP_CONFIG.JWTConfig.SecretKey),
		ExpiresTime: app.AppContext.APP_CONFIG.JWTConfig.ExpiresTime,
		BufferTime:  app.AppContext.APP_CONFIG.JWTConfig.BufferTime,
		Issuer:      app.AppContext.APP_CONFIG.JWTConfig.Issuer,
	}
}

func (jwtUtil *JWTUtil) CreateClaims(mobile string, userName string) apiCommon.JWTClaims {
	claims := apiCommon.JWTClaims{
		Mobile:     mobile,
		UserName:   userName,
		BufferTime: jwtUtil.BufferTime, //缓冲时间内会获得新的token刷新令牌
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        GenerateId(),
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                                // 签名生效时间
			NotBefore: jwt.NewNumericDate(time.Unix(time.Now().Unix()-1000, 0).UTC()),                // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Unix(time.Now().Unix()+jwtUtil.ExpiresTime, 0).UTC()), // 过期时间
			Issuer:    jwtUtil.Issuer,                                                                // 签名的发行者
		},
	}
	return claims
}

func (jwtUtil *JWTUtil) GenerateToken(claims apiCommon.JWTClaims) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtUtil.SecretKey)
	if err != nil {
		fmt.Println(err)
	}
	return token, err
}

func (jwtUtil *JWTUtil) Refresh(oldToken string, claims apiCommon.JWTClaims) (string, error) {
	v, err, _ := app.AppContext.APP_Concurrency_Controller.Do("JWT:"+oldToken, func() (interface{}, error) {
		c, perr := jwtUtil.ParseWithClaims(oldToken) 
		if c == nil {
			return nil, perr
		}
		issuedAt, _ := c.RegisteredClaims.GetExpirationTime()
		expiredAt := time.Now()
		if issuedAt != nil {
			expiredAt = time.Unix(issuedAt.Unix()+c.BufferTime, 0).UTC()
		}

		if time.Now().After(expiredAt)) {
			return nil, TokenExpired
		}
		return jwtUtil.GenerateToken(claims)
	})
	return v.(string), err
}

func (jwtUtil *JWTUtil) ParseToken(tokenString string) (*apiCommon.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &apiCommon.JWTClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return jwtUtil.SecretKey, nil
	})
	if err != nil {
		switch err {
		case jwt.ErrTokenMalformed:
			return nil, TokenMalformed
		case jwt.ErrTokenExpired:
			return nil, TokenExpired
		case jwt.ErrTokenNotValidYet:
			return nil, TokenNotValid
		default:
			return nil, TokenInvalid
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*apiCommon.JWTClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}
