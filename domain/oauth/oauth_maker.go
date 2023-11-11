package oauth

type OAuthMaker interface {
	GenerateAccessToken(mobile string, username string) (string, *OAuthClaims, error)
	GenerateRefreshToken(mobile string, username string) (string, *OAuthClaims, error)
	VerifyToken(token string) (*OAuthClaims, error)
}
