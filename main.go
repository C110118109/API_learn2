package main

import (
	//todo 路徑打錯ㄌ

	"fmt"
	"net/http"

	_ "eirc.app/api"
	"eirc.app/internal/pkg/dao/gorm"
	"eirc.app/internal/pkg/log"

	"eirc.app/internal/v1/router"
	"eirc.app/internal/v1/router/department"
	"eirc.app/internal/v1/router/employee"
	"eirc.app/internal/v1/router/item"
	"eirc.app/internal/v1/router/request"
	"eirc.app/internal/v1/router/request_itemlist"
	"eirc.app/internal/v1/router/role"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	db, err := gorm.New()

	if err != nil {

		log.Error(err)
		return
	}

	route := router.Default()
	//todo 這裡也是
	route = request.GetRoute(route, db)
	route = department.GetRoute(route, db)
	route = item.GetRoute(route, db)
	route = employee.GetRoute(route, db)
	route = role.GetRoute(route, db)
	route = request_itemlist.GetRoute(route, db)

	url := ginSwagger.URL(fmt.Sprintf("http://localhost:8080/swagger/doc.json"))
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	log.Fatal(http.ListenAndServe(":8080", route))
}
