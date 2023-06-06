package request

import (
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/request"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("request")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":requestID", controller.GetByID)
		v10.GET("/GetBysingle/:requestID", controller.GetBySingle)
		v10.GET("/GetByReIDDetail/:requestID", controller.GetByReIDRequestDetailListUser)
		v10.GET("/GetAllRequestDetail", controller.RequestDetailListUser)
		v10.DELETE(":requestID", controller.Delete)
		v10.PATCH(":requestID", controller.Updated)
	}

	return route
}
