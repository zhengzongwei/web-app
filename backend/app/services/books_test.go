// Package          services
// @Title           books_test.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/11/16 16:56

package services

import (
	"backend/app/database"
	"backend/app/models"
	"testing"
)

func TestBookCreate(t *testing.T) {
	bookService := DBBookService{DB: database.GetDB()}
	books := []models.Book{
		{
			Name:    "C#程序设计244",
			Comment: "测试123",
			Authors: []*models.Author{
				{
					Name: "作者测试1",
				},
				{
					Name: "作者测试2",
				},
				{
					Name: "作者测试3",
				},
			},
		},
		{
			Name:    "C#程序设计2t3",
			Comment: "测试",
			Authors: []*models.Author{
				{
					Name: "作者测试1",
				},
				{
					Name: "作者测试2",
				},
				{
					Name: "作者测试6",
				},
			},
		},
	}
	err := bookService.BookCreate(&books)
	if err != nil {
		return
	}
}

//	result, err := BookCreate(&book)
//	if err != nil {
//		return
//	}
//	fmt.Printf("%d \n", result)
//}
//
//func TestBookList(t *testing.T) {
//	books := BookList()
//	fmt.Printf("111 %v \n", books)
//}
//
//func TestBookDetail(t *testing.T) {
//	bookId := 9
//	book := BookDetail(uint(bookId))
//	fmt.Printf("111 %v \n", book)
//
//}

//func TestBookDelete(t *testing.T) {
//	bookIds := []uint{9, 10}
//	BookDelete(bookIds)
//
//}
