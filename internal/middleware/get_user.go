package middleware

import (
	"getcare-notification/internal/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tableName := service.GetTableNameFromPath(ctx.FullPath())
		if CheckWhitelist(tableName) {
			return
		}

		token := ctx.GetHeader("token")
		if token == "" {
			return
		}

		if token == os.Getenv("TOKEN_SERVER1") ||
			token == os.Getenv("TOKEN_SERVER2") ||
			token == os.Getenv("TOKEN_DEFAULT") {
			return
		}
		userInfo, err := service.GetKeyCloakUserInfo(&gin.Context{}, token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"result":     false,
				"message":    "Not Valid Token",
				"error_code": 401,
			})
			return
		}
		ctx.Set("user", userInfo)
	}
}
