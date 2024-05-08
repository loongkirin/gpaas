package core

import (
	jwt "github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	jwt.RegisteredClaims
	Mobile     string
	UserName   string
	BufferTime int64
}
