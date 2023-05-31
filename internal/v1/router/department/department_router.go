package department

import (
	"eirc.app/internal/v1/middleware"
	"eirc.app/internal/v1/presenter/department"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := department.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("departments")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":departmentID", controller.GetByID)
		v10.DELETE(":departmentID", controller.Delete)
		v10.PATCH(":departmentID", controller.Updated)
	}

	return route
}
