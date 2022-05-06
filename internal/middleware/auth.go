package middleware

import (
	"getcare-notification/internal/service"
	"getcare-notification/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var WhitelistTable = []string{}

func Authorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tableName := service.GetTableNameFromPath(ctx.FullPath())
		if CheckWhitelist(tableName) {
			return
		}

		token := ctx.GetHeader("token")
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"result":     false,
				"message":    "No Authorization header found",
				"error_code": 401,
			})
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

func CheckWhitelist(tableName string) bool {
	return utils.SliceContains(WhitelistTable, tableName)
}
