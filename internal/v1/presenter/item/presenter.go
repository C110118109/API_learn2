package item

import (
	"eirc.app/internal/v1/resolver/item"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Presenter interface {
	Created(ctx *gin.Context)
	List(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Updated(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type presenter struct {
	ItemResolver item.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		ItemResolver: item.New(db),
	}
}
