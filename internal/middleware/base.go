package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"getcare-notification/internal/service"
	"getcare-notification/utils"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

func Base() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		Log(ctx)
		ctx.Next()
	}
}

func Log(ctx *gin.Context) {
	var bodyBytes []byte
	if ctx.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(ctx.Request.Body)
	}
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	query := ctx.Request.URL.Query()
	paramJson, err := json.Marshal(query)
	if err != nil {
		paramJson = []byte{}
	}

	clientIP := ctx.ClientIP()
	method := ctx.Request.Method
	fullPath := ctx.FullPath()
	statusCode := fmt.Sprint(ctx.Writer.Status())
	proto := ctx.Request.Proto
	userAgent := ctx.Request.UserAgent()
	bodyStr := string(bodyBytes)
	paramStr := string(paramJson)

	log.Println(fmt.Sprintf("[%s] %s %s %s %s \"%s\" %s \n%s",
		clientIP,
		method,
		fullPath,
		statusCode,
		proto,
		userAgent,
		paramStr,
		bodyStr,
	))

	logID := service.GeneratorString(32)
	ctx.Set(utils.REQUEST_LOG_ID, logID)
}
