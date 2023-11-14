package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Engin
	router := gin.Default()
	//router := gin.New()

	router.GET("/", func(context *gin.Context) {
		log.Println(">>>> hello gin start <<<<")
		context.JSON(200, gin.H{
			"code":    200,
			"success": true,
		})
	})
	// 指定地址和端口号
	err := router.Run("localhost:8080")
	if err != nil {
		log.Println(">>>> ERROR <<<<")
	}
}
