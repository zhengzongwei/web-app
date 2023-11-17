// Package          services
// @Title           books_test.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/11/16 16:56

package services

import (
	"backend/app/models"
	"fmt"
	"testing"
)

func TestBookCreate(t *testing.T) {
	book := models.Book{
		Name:    "C#程序设计1",
		Comment: "测试123",
	}

	result, err := BookCreate(&book)
	if err != nil {
		return
	}
	fmt.Printf("%d \n", result)
}

func TestBookList(t *testing.T) {
	books := BookList()
	fmt.Printf("111 %v \n", books)
}

func TestBookDetail(t *testing.T) {
	bookId := 9
	book := BookDetail(uint(bookId))
	fmt.Printf("111 %v \n", book)

}

func TestBookDelete(t *testing.T) {
	bookIds := []uint{9, 10}
	BookDelete(bookIds)

}
