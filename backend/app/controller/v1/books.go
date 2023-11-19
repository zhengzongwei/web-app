// Package          v1
// @Title           books.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/11/16 16:04

package v1

import (
	"backend/app/database"
	"backend/app/models"
	"backend/app/response"
	"backend/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddBooks(c *gin.Context) {

	//var authors models.AuthorList
	var books models.BookList

	if err := c.BindJSON(&books); err != nil {
		response.Response(c, http.StatusBadRequest, -1, gin.H{"err": err.Error()})
		return
	}

	bookService := &services.DBBookService{DB: database.GetDB()}

	err := bookService.BookCreate(&books.Books)
	if err != nil {
		response.Response(c, http.StatusBadRequest, -1, gin.H{"err": err})
	}
}

func ListBooks(c *gin.Context) {
	bookService := &services.DBBookService{DB: database.GetDB()}
	books := bookService.BookList()
	response.Response(c, http.StatusOK, 0, gin.H{"books": books})
	return
}

func DetailBook(c *gin.Context) {

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	bookService := &services.DBBookService{DB: database.GetDB()}
	book := bookService.BookDetail(uint(id))
	response.Response(c, http.StatusOK, 0, gin.H{"book": book})
	return
}

func DeleteBooks(c *gin.Context) {

	var books models.DelBookList
	bookService := &services.DBBookService{DB: database.GetDB()}
	if err := c.BindJSON(&books); err != nil {
		response.Response(c, http.StatusBadRequest, -1, gin.H{"err": err})
		return
	}

	_, err := bookService.BookDelete(&books.BookIds)
	if err != nil {
		return
	}

	response.Response(c, http.StatusOK, 0, nil)
}
