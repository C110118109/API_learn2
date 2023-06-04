package request_itemlist

import (
	"net/http"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/request_itemlists"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (p *presenter) Created(ctx *gin.Context) {

	trx := ctx.MustGet("db_trx").(*gorm.DB)
	createBy := util.GenerateUUID()
	input := &request_itemlists.Created{}
	input.CreatedBy = createBy
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.RequestItemListResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

func (p *presenter) List(ctx *gin.Context) {
	input := &request_itemlists.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.RequestItemListResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

func (p *presenter) GetBySingle(ctx *gin.Context) {
	requestitemlistID := ctx.Param("requestitemlistID")
	input := &request_itemlists.Base{}
	input.RequestItemListID = requestitemlistID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.RequestItemListResolver.GetBySingle(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

func (p *presenter) GetByID(ctx *gin.Context) {
	requestitemlistID := ctx.Param("requestitemlistID")
	input := &request_itemlists.Field{}
	input.RequestItemListID = requestitemlistID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.RequestItemListResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

func (p *presenter) Delete(ctx *gin.Context) {

	requestitemlistID := ctx.Param("requestitemlistID")
	input := &request_itemlists.Updated{}
	input.RequestItemListID = requestitemlistID

	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.RequestItemListResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

func (p *presenter) Updated(ctx *gin.Context) {

	requestitemlistID := ctx.Param("requestitemlistID")
	input := &request_itemlists.Updated{}
	input.RequestItemListID = requestitemlistID

	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.RequestItemListResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
