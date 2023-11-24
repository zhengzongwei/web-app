// Package          v1
// @Title           books.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/11/16 16:04

package books

import (
	"backend/app/database"
	"backend/app/models"
	"backend/app/response"
	"backend/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateBook(c *gin.Context) {
	var createBooks struct {
		Books []models.Book `json:"books"`
	}

	// 从请求中解析 JSON 数据并绑定到 Book 结构体
	if err := c.BindJSON(&createBooks); err != nil {
		response.Response(c, http.StatusInternalServerError, -1, gin.H{"error": err.Error()})
		return
	}

	// 调用 Service 层进行书籍创建操作
	bookService := services.NewBookService(database.GetDB())
	err := bookService.CreateBook(createBooks.Books)

	// 处理错误并返回响应
	if err != nil {
		response.Response(c, http.StatusInternalServerError, -1, gin.H{"error": err.Error()})
		return
	}
	response.Response(c, http.StatusOK, 0, gin.H{})
}

func ListBook(c *gin.Context) {
	bookService := services.NewBookService(database.GetDB())

	// 调用 Service 层的 ListBook 函数获取书籍列表
	books, err := bookService.ListBook()

	// 处理错误并返回适当的 HTTP 响应
	if err != nil {
		response.Response(c, http.StatusInternalServerError, -1, gin.H{"error": err.Error()})
		return
	}

	// 返回成功的 HTTP 响应
	response.Response(c, http.StatusOK, 0, gin.H{"data": books})
}

func DeleteBook(c *gin.Context) {
	// 从请求中解析 JSON 数据并绑定到结构体
	var request struct {
		BookIds []uint `json:"book_ids"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		response.Response(c, http.StatusBadRequest, -1, gin.H{"error": err.Error()})
		return
	}

	// 调用 Service 层进行批量删除操作
	bookService := services.NewBookService(database.GetDB())
	if err := bookService.DeleteBook(request.BookIds); err != nil {
		response.Response(c, http.StatusInternalServerError, -1, gin.H{"error": err.Error()})

	}
	// 返回成功的 HTTP 响应
	response.Response(c, http.StatusOK, 0, gin.H{})
}

func EditBook(c *gin.Context) {
	var book models.Book
	bookID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Response(c, http.StatusBadRequest, -1, gin.H{"error": "Invalid book ID"})
		return
	}

	if err := c.Bind(&book); err != nil {
		response.Response(c, http.StatusBadRequest, -1, gin.H{"error": err.Error()})
		return
	}
	bookService := services.NewBookService(database.GetDB())
	if err := bookService.EditBook(uint(bookID), &book); err != nil {
		response.Response(c, http.StatusInternalServerError, -1, gin.H{"error": err.Error()})
	}
	// 返回成功的 HTTP 响应
	response.Response(c, http.StatusOK, 0, gin.H{})

}
func DetailBook(c *gin.Context) {

	// 从 URL 参数中获取书籍 ID
	bookID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Response(c, http.StatusBadRequest, -1, gin.H{"error": "Invalid book ID"})
		return
	}

	// 调用 Service 层获取书籍详情
	bookService := services.NewBookService(database.GetDB())
	book, err := bookService.DetailBook(uint(bookID))

	// 处理错误并返回适当的 HTTP 响应
	if err != nil {
		response.Response(c, http.StatusInternalServerError, -1, gin.H{"error": err.Error()})
		return
	}

	// 返回成功的 HTTP 响应
	response.Response(c, http.StatusOK, 0, gin.H{"book": book})
}
