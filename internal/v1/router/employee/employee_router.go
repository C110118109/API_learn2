package employee

import (
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/employee"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("employee")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":employeeID", controller.GetByID)
		v10.GET("/GetBysingle/:employeeID", controller.GetBySingle)
		v10.DELETE(":employeeID", controller.Delete)
		v10.PATCH(":employeeID", controller.Updated)
	}

	return route
}
