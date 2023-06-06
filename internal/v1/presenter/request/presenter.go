package request

import (
	"eirc.app/internal/v1/resolver/request"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Presenter interface {
	Created(ctx *gin.Context)
	List(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	GetBySingle(ctx *gin.Context)
	GetByReIDRequestDetailListUser(ctx *gin.Context)
	RequestDetailListUser(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Updated(ctx *gin.Context)
}

type presenter struct {
	RequestResolver request.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		RequestResolver: request.New(db),
	}
}
