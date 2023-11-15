// Package          router
// @Title           router.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/11/15 16:16

package v1

import (
	controller "backend/app/controller/v1"
	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	router := gin.Default()
	router.GET("/", controller.IndexApi)
	return router
}
