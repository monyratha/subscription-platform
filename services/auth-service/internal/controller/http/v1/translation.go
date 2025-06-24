package v1

import (
	"net/http"

	"github.com/evrone/go-clean-template/internal/controller/http/v1/request"
	"github.com/evrone/go-clean-template/internal/entity"
	"github.com/gin-gonic/gin"
)

// @Summary     Show history
// @Description Show all translation history
// @ID          history
// @Tags  	    translation
// @Accept      json
// @Produce     json
// @Success     200 {object} entity.TranslationHistory
// @Failure     500 {object} response.Error
// @Router      /translation/history [get]
func (r *V1) history(ctx *gin.Context) {
	translationHistory, err := r.t.History(ctx.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - history")
		errorResponse(ctx, http.StatusInternalServerError, "database problems")
		return
	}

	ctx.JSON(http.StatusOK, translationHistory)
}

// @Summary     Translate
// @Description Translate a text
// @ID          do-translate
// @Tags  	    translation
// @Accept      json
// @Produce     json
// @Param       request body request.Translate true "Set up translation"
// @Success     200 {object} entity.Translation
// @Failure     400 {object} response.Error
// @Failure     500 {object} response.Error
// @Router      /translation/do-translate [post]
func (r *V1) doTranslate(ctx *gin.Context) {
	var body request.Translate

	if err := ctx.ShouldBindJSON(&body); err != nil {
		r.l.Error(err, "http - v1 - doTranslate")
		errorResponse(ctx, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := r.v.Struct(body); err != nil {
		r.l.Error(err, "http - v1 - doTranslate")
		errorResponse(ctx, http.StatusBadRequest, "invalid request body")
		return
	}

	translation, err := r.t.Translate(
		ctx.Request.Context(),
		entity.Translation{
			Source:      body.Source,
			Destination: body.Destination,
			Original:    body.Original,
		},
	)
	if err != nil {
		r.l.Error(err, "http - v1 - doTranslate")
		errorResponse(ctx, http.StatusInternalServerError, "translation service problems")
		return
	}

	ctx.JSON(http.StatusOK, translation)
}
