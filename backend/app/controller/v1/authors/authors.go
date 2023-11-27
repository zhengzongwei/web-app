// Package          authors
// @Title           authors.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/11/27 10:32

package authors

import (
	"backend/app/database"
	"backend/app/models"
	"backend/app/response"
	"backend/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateAuthor(c *gin.Context) {
	var createAuthor struct {
		Authors []models.Author `json:"authors"`
	}

	if err := c.BindJSON(&createAuthor); err != nil {
		response.Response(c, http.StatusInternalServerError, -1, gin.H{"error": err.Error()})
		return
	}
	authorService := services.NewAuthorService(database.GetDB())
	err := authorService.CreateAuthor(createAuthor.Authors)
	// 处理错误并返回响应
	if err != nil {
		response.Response(c, http.StatusInternalServerError, -1, gin.H{"error": err.Error()})
		return
	}
	response.Response(c, http.StatusOK, 0, gin.H{})
}

func ListAuthor(c *gin.Context) {
	authorService := services.NewAuthorService(database.GetDB())
	authors, err := authorService.ListAuthor()
	if err != nil {
		response.Response(c, http.StatusInternalServerError, -1, gin.H{"error": err.Error()})
		return
	}
	response.Response(c, http.StatusOK, 0, gin.H{"authors": authors})

}

func DeleteAuthor(c *gin.Context) {
	var request struct {
		AuthorIDs []uint `json:"author_ids"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Response(c, http.StatusBadRequest, -1, gin.H{"error": err.Error()})
		return
	}
	authorService := services.NewAuthorService(database.GetDB())
	if err := authorService.DeleteAuthor(request.AuthorIDs); err != nil {
		response.Response(c, http.StatusInternalServerError, -1, gin.H{"error": err.Error()})
		return
	}
	// 返回成功的 HTTP 响应
	response.Response(c, http.StatusOK, 0, gin.H{})
}

func EditAuthor(c *gin.Context) {
	var author models.Author
	authorID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Response(c, http.StatusBadRequest, -1, gin.H{"error": "Invalid book ID"})
		return
	}
	if err := c.Bind(&author); err != nil {
		response.Response(c, http.StatusBadRequest, -1, gin.H{"error": err.Error()})
		return
	}
	authorService := services.NewAuthorService(database.GetDB())
	if err := authorService.EditAuthor(uint(authorID), &author); err != nil {
		response.Response(c, http.StatusInternalServerError, -1, gin.H{"error": err.Error()})
	}
	// 返回成功的 HTTP 响应
	response.Response(c, http.StatusOK, 0, gin.H{})

}

func DetailAuthor(c *gin.Context) {
	// 从 URL 参数中获取书籍 ID
	authorID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Response(c, http.StatusBadRequest, -1, gin.H{"error": "Invalid author ID"})
		return
	}
	authorService := services.NewAuthorService(database.GetDB())
	author, err := authorService.DetailAuthor(uint(authorID))
	if err != nil {
		response.Response(c, http.StatusInternalServerError, -1, gin.H{"error": err.Error()})
		return
	}
	response.Response(c, http.StatusOK, 0, gin.H{"author": author})

}
