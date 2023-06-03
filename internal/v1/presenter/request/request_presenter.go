package request

import (
	"net/http"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/requests"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (p *presenter) Created(ctx *gin.Context) {

	trx := ctx.MustGet("db_trx").(*gorm.DB)
	createBy := util.GenerateUUID()
	input := &requests.Created{}
	input.CreatedBy = createBy
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.RequestResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

func (p *presenter) List(ctx *gin.Context) {
	input := &requests.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.RequestResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

func (p *presenter) GetBySingle(ctx *gin.Context) {
	requestID := ctx.Param("requestID")
	input := &requests.Base{}
	input.RequestID = requestID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.RequestResolver.GetBySingle(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

func (p *presenter) GetByID(ctx *gin.Context) {
	requestID := ctx.Param("requestID")
	input := &requests.Field{}
	input.RequestID = requestID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.RequestResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

func (p *presenter) Delete(ctx *gin.Context) {

	requestID := ctx.Param("requestID")
	input := &requests.Updated{}
	input.RequestID = requestID

	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.RequestResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

func (p *presenter) Updated(ctx *gin.Context) {

	requestID := ctx.Param("rrequestID")
	input := &requests.Updated{}
	input.RequestID = requestID

	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.RequestResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
