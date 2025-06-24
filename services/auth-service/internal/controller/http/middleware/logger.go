package middleware

import (
	"strconv"
	"strings"

	"github.com/evrone/go-clean-template/pkg/logger"
	"github.com/gin-gonic/gin"
)

func buildRequestMessage(ctx *gin.Context) string {
	var result strings.Builder

	result.WriteString(ctx.ClientIP())
	result.WriteString(" - ")
	result.WriteString(ctx.Request.Method)
	result.WriteString(" ")
	result.WriteString(ctx.Request.URL.String())
	result.WriteString(" - ")
	result.WriteString(strconv.Itoa(ctx.Writer.Status()))
	result.WriteString(" ")
	result.WriteString(strconv.Itoa(ctx.Writer.Size()))

	return result.String()
}

func Logger(l logger.Interface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		l.Info(buildRequestMessage(ctx))
	}
}
