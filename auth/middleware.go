package auth

import (
	"crypto/subtle"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireBearerToken(token string) gin.HandlerFunc {
	expected := "Bearer " + token

	return func(ctx *gin.Context) {
		if subtle.ConstantTimeCompare([]byte(ctx.GetHeader("Authorization")), []byte(expected)) != 1 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "missing or invalid authorization token",
			})
			return
		}

		ctx.Next()
	}
}
