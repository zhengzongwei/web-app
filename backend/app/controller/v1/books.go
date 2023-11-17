// Package          v1
// @Title           books.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/11/16 16:04

package v1

import (
	"backend/app/response"
	"backend/app/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//type Books struct {
//	Name    string `json:"name"`
//	Comment string `json:"comment"`
//}

func AddBooks(c *gin.Context) {

	var book services.BookData
	if err := c.BindJSON(&book); err != nil {
		response.Response(c, http.StatusBadRequest, -1, nil)
	}
	id, err := services.Create(&book)
	if err != nil {
		return
	}
	fmt.Printf("%s", book)
	response.Response(c, http.StatusOK, 0, gin.H{"id": id})
}
