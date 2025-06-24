package v1

import (
	"github.com/evrone/go-clean-template/internal/controller/http/v1/response"
	"github.com/gin-gonic/gin"
)

func errorResponse(ctx *gin.Context, code int, msg string) {
	ctx.JSON(code, response.Error{Error: msg})
}
