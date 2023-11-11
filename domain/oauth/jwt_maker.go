package oauth

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	core "github.com/loongkirin/gpaas/core"
)

const MinSecretKeySize = 32

// JWTMaker is a JSON Web Token maker.
type JWTMaker struct {
	oauthConfig core.OAuthConfig
}

func (maker JWTMaker) GenerateToken(mobile string, username string) (string, *OAuthClaims, error) {
	claims := NewOAuthClaims(mobile, username, maker.oauthCfg.Issuer, maker.oauthCfg.ExpiresTime)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := jwtToken.SignedString([]byte(maker.oauthCfg.SecretKey))
	return token, claims, err
}

func (maker JWTMaker) VerifyToken(token string) (*OAuthClaims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidKey
		}
		return []byte(maker.oauthCfg.SecretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &OAuthClaims{}, keyFunc)
	if err != nil {
		// TODO check if we really need that.
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrTokenExpired
		}

		return nil, err
	}

	claims, ok := jwtToken.Claims.(*OAuthClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

// NewJWTMaker creates a new JWTMaker.
func NewJWTMaker(oauthCfg core.OAuthConfig) (Maker, error) {
	if len(oauthCfg.SecretKey) < MinSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", MinSecretKeySize)
	}

	return &JWTMaker{oauthCfg: oauthCfg}, nil
}
