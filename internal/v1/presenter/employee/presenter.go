package employee

import (
	"eirc.app/internal/v1/resolver/employee"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Presenter interface {
	Created(ctx *gin.Context)
	List(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Updated(ctx *gin.Context)
}

type presenter struct {
	Employeeesolver employee.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		Employeeesolver: employee.New(db),
	}
}
