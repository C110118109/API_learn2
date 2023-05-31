package item

import (
	"eirc.app/internal/v1/middleware"
	"eirc.app/internal/v1/presenter/item"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := item.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("items")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":itemID", controller.GetByID)
		v10.DELETE(":itemID", controller.Delete)
		v10.PATCH(":itemID", controller.Updated)
	}

	return route
}
