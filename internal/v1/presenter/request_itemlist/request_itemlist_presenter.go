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

// GiftDetailListUser
// @Summary 取得全部部品零件贈送申請單及明細
// @description 取得全部部品零件贈送申請單及明細
// @Tags GiftApplication
// @version 1.0
// @Accept json
// @produce json
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=gift_applications.AllGiftDetail} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/GiftApplication/GetAllGiftApplication [get]
func (p *presenter) ItemDetailUser(ctx *gin.Context) {
	input := &request_itemlists.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.RequestItemListResolver.ItemDetailUser(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByGIDGiftDetailListUser
// @Summary 取得單一部品零件贈送申請單及明細
// @description 取得單一部品零件贈送申請單及明細
// @Tags GiftApplication
// @version 1.0
// @Accept json
// @produce json
// @param fID path string true "部品零件贈送申請單ID"
// @success 200 object code.SuccessfulMessage{body=gift_applications.GiftDetail} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/GiftApplication/GetByGIDDetail/{gID} [get]
func (p *presenter) GetByIIDItemDetailUser(ctx *gin.Context) {
	requestitemlistID := ctx.Param("requestitemlistID")
	input := &request_itemlists.Field{}
	input.RequestItemListID = requestitemlistID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.RequestItemListResolver.GetByIIDItemDetailUser(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
