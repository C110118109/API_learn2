package request_itemlist

import (
	"eirc.app/internal/v1/resolver/request_itemlist"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Presenter interface {
	Created(ctx *gin.Context)
	List(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	GetBySingle(ctx *gin.Context)
	GetByIIDItemDetailUser(ctx *gin.Context)
	ItemDetailUser(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Updated(ctx *gin.Context)
}

type presenter struct {
	RequestItemListResolver request_itemlist.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		RequestItemListResolver: request_itemlist.New(db),
	}
}
