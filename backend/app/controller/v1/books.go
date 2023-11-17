// Package          v1
// @Title           books.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/11/16 16:04

package v1

import (
	"backend/app/models"
	"backend/app/response"
	"backend/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//type Books struct {
//	Name    string `json:"name"`
//	Comment string `json:"comment"`
//}

func AddBooks(c *gin.Context) {

	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		response.Response(c, http.StatusBadRequest, -1, nil)
		return
	}

	id, err := services.BookCreate(&book)
	if err != nil {
		return
	}

	response.Response(c, http.StatusOK, 0, gin.H{"id": id})
	return
}

func ListBooks(c *gin.Context) {
	books := services.BookList()
	response.Response(c, http.StatusOK, 0, gin.H{"books": books})
	return
}

func DetailBook(c *gin.Context) {

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	book := services.BookDetail(uint(id))
	response.Response(c, http.StatusOK, 0, gin.H{"book": book})
	return
}

func DeleteBooks(c *gin.Context) {

	var books models.Books

	if err := c.ShouldBindJSON(&books); err != nil {
		response.Response(c, http.StatusBadRequest, -1, gin.H{"err": err})
		return
	}

	for _, book := range books.BookIds {
		_, err := services.BookDelete(&book)
		if err != nil {
			return
		}
	}

	response.Response(c, http.StatusOK, 0, nil)
}
