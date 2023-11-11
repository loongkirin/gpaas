package oauth

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	util "github.com/loongkirin/gpaas/util"
)

var (
	ErrInvalidKey   = errors.New("key is invalid")
	ErrTokenExpired = errors.New("Token is expired")
	ErrTokenInvalid = errors.New("Tlken is invalid")
)

type OAuthClaims struct {
	Id        string    `json:"id"`
	Mobile    string    `json:"mobile"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
	NotBefore time.Time `json:"not_before"`
	Issuer    string    `json:"issuer,omitempty"`
}

func (o OAuthClaims) GetExpirationTime() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(o.ExpiredAt), nil
}

func (o OAuthClaims) GetIssuedAt() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(o.IssuedAt), nil
}

func (o OAuthClaims) GetNotBefore() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(o.NotBefore), nil
}

func (o OAuthClaims) GetIssuer() (string, error) {
	return o.Issuer, nil
}

func (o OAuthClaims) GetSubject() (string, error) {
	return "subject", nil
}

func (o OAuthClaims) GetAudience() (jwt.ClaimStrings, error) {
	return jwt.ClaimStrings{"audience"}, nil
}

func (o OAuthClaims) Valid() error {
	if o.ExpiredAt.Before(time.Now()) {
		return ErrTokenExpired
	}

	return nil
}

func NewOAuthClaims(mobile string, username string, issuer string, duration int64) *OAuthClaims {
	claims := &OAuthClaims{
		Id:        util.GenerateId(),
		Mobile:    mobile,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: jwt.NewNumericDate(time.Unix(time.Now().Unix()+duration, 0).UTC()),
		NotBefore: jwt.NewNumericDate(time.Unix(time.Now().Unix()-1000, 0).UTC()),
		Issuer:    issuer,
	}

	return claims
}
