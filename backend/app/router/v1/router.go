// Package          router
// @Title           router.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/11/15 16:16

package v1

import (
	controller "backend/app/controller/v1"
	"backend/app/controller/v1/authors"
	"backend/app/controller/v1/books"
	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	router := gin.Default()
	router.GET("/", controller.IndexApi)
	api := router.Group("api")
	{
		v1 := api.Group("v1")
		book := v1.Group("book")
		{
			book.GET("list", books.ListBook)
			book.GET("detail/:id", books.DetailBook)
			book.POST("create", books.CreateBook)
			book.DELETE("delete", books.DeleteBook)
			book.PUT("edit/:id", books.EditBook)
		}
		author := v1.Group("author")
		{
			author.GET("list", authors.ListAuthor)
			author.POST("create", authors.CreateAuthor)

		}

	}

	return router
}
