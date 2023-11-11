package oauth

import (
	"fmt"

	core "github.com/loongkirin/gpaas/core"
	"github.com/o1egl/paseto/v2"
	"golang.org/x/crypto/chacha20poly1305"
)

// PasetoMaker is a PASETO maker.
type PasetoMaker struct {
	paseto      *paseto.V2
	secretKey   []byte
	oauthConfig core.OAuthConfig
}

func (maker PasetoMaker) GenerateToken(mobile string, username string) (string, *OAuthClaims, error) {
	claims := NewOAuthClaims(mobile, username, maker.oauthCfg.Issuer, maker.oauthCfg.ExpiresTime)
	token, err := maker.paseto.Encrypt(maker.secretKey, claims, nil)
	return token, claims, err
}

func (maker PasetoMaker) VerifyToken(token string) (*OAuthClaims, error) {
	claims := &OAuthClaims{}
	err := maker.paseto.Decrypt(token, p.secretKey, claims, nil)
	if err != nil {
		return nil, ErrTokenInvalid
	}

	err = claims.Valid()
	if err != nil {
		return nil, err
	}

	return claims, nil
}

// NewPasetoMaker creates a new PasetoMaker.
func NewPasetoMaker(oauthCfg core.OAuthConfig) (OAuthMaker, error) {
	if len(oauthCfg.SecretKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:    paseto.NewV2(),
		secretKey: []byte(oauthCfg.SecretKey),
		oauthCfg:  oauthCfg,
	}

	return maker, nil
}
