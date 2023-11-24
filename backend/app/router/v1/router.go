// Package          router
// @Title           router.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/11/15 16:16

package v1

import (
	controller "backend/app/controller/v1"
	"backend/app/controller/v1/books"
	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	router := gin.Default()
	router.GET("/", controller.IndexApi)
	api := router.Group("api")
	{
		v1 := api.Group("v1")
		{
			v1.GET("listbook", books.ListBook)
			v1.GET("detailbook/:id", books.DetailBook)
			v1.POST("addbook", books.CreateBook)
			v1.DELETE("delbook", books.DeleteBook)
			v1.PUT("editbook/:id", books.EditBook)

		}

	}

	return router
}
