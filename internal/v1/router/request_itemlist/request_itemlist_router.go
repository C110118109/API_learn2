package request_itemlist

import (
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/request_itemlist"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("request_itemlist")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":request_itemlistID", controller.GetByID)
		v10.DELETE(":request_itemlistID", controller.Delete)
		v10.PATCH(":request_itemlistID", controller.Updated)
	}

	return route
}
