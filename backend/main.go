package main

import (
	"backend/app/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// Engin
	router := gin.Default()
	//router := gin.New()

	router.GET("/", func(context *gin.Context) {
		log.Println(">>>> hello gin start <<<<")
		response.Response(context, http.StatusOK, 0, nil)
	})
	// 指定地址和端口号
	err := router.Run("localhost:8080")
	if err != nil {
		log.Println(">>>> ERROR <<<<")
	}
}
