package oauth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	cfg "github.com/loongkirin/gpaas/pkg/config"
)

const MinSecretKeySize = 32

// JWTMaker is a JSON Web Token maker.
type JWTMaker struct {
	oauthConfig cfg.OAuthConfig
}

func (maker JWTMaker) GenerateAccessToken(mobile string, username string) (string, *OAuthClaims, error) {
	duration, err := time.ParseDuration(maker.oauthConfig.AccessExpiresTime)
	if err != nil {
		duration = time.Hour * 8
	}

	claims := NewOAuthClaims(mobile, username, maker.oauthConfig.Issuer, duration)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(maker.oauthConfig.SecretKey))
	return token, claims, err
}

func (maker JWTMaker) GenerateRefreshToken(mobile string, username string) (string, *OAuthClaims, error) {
	duration, err := time.ParseDuration(maker.oauthConfig.RefreshExpiresTime)
	if err != nil {
		duration = time.Hour * 24 * 7
	}

	claims := NewOAuthClaims(mobile, username, maker.oauthConfig.Issuer, duration)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(maker.oauthConfig.SecretKey))
	return token, claims, err
}

func (maker JWTMaker) VerifyToken(token string) (*OAuthClaims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidKey
		}
		return []byte(maker.oauthConfig.SecretKey), nil
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
		return nil, ErrTokenInvalid
	}

	return claims, nil
}

// NewJWTMaker creates a new JWTMaker.
func NewJWTMaker(oauthCfg cfg.OAuthConfig) (OAuthMaker, error) {
	if len(oauthCfg.SecretKey) < MinSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", MinSecretKeySize)
	}

	return &JWTMaker{oauthConfig: oauthCfg}, nil
}
