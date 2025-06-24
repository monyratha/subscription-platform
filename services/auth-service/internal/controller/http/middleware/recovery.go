package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/evrone/go-clean-template/pkg/logger"
	"github.com/gin-gonic/gin"
)

func buildPanicMessage(ctx *gin.Context, err interface{}) string {
	var result strings.Builder

	result.WriteString(ctx.ClientIP())
	result.WriteString(" - ")
	result.WriteString(ctx.Request.Method)
	result.WriteString(" ")
	result.WriteString(ctx.Request.URL.String())
	result.WriteString(" PANIC DETECTED: ")
	result.WriteString(fmt.Sprintf("%v\n%s\n", err, debug.Stack()))

	return result.String()
}

func logPanic(l logger.Interface) func(c *gin.Context, err interface{}) {
	return func(ctx *gin.Context, err interface{}) {
		l.Error(buildPanicMessage(ctx, err))
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}
}

func Recovery(l logger.Interface) gin.HandlerFunc {
	return gin.CustomRecovery(logPanic(l))
}
