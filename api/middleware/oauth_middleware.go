package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	response "github.com/loongkirin/gpaas/api/common"
	oauth "github.com/loongkirin/gpaas/domain/oauth"
)

const (
	authorizationHeaderKey  = "x-auth"
	authorizationTypeBearer = "x-bearer"
	authorizationClaimsKey  = "x-auth-claims"
)

func OAuthMiddleware(oauthMaker oauth.OAuthMaker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.NewResponse(response.UNAUTHORIZED, "Authorization Header Invalid"))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.NewResponse(response.UNAUTHORIZED, "Authorization Header Invalid"))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.NewResponse(response.UNAUTHORIZED, "Authorization Type Invalid"))
			return
		}

		accessToken := fields[1]
		claims, err := oauthMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.NewResponse(response.UNAUTHORIZED, err.Error()))
			return
		}

		ctx.Set(authorizationClaimsKey, claims)
		ctx.Next()
	}
}
