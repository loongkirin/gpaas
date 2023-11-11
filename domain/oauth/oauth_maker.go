package oauth

type OAuthMaker interface {
	GenerateToken(mobile string, username string) (string, *OAuthClaims, error)
	VerifyToken(token string) (*OAuthClaims, error)
}
