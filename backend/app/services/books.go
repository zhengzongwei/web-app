// Package          services
// @Title           books.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/11/16 15:49

package services

import (
	"backend/app/database"
	"backend/app/models"
)

type BookData struct {
	// 书名
	Name string
	//Authors   []byte
	//Publishes []byte
	Comment string
}

//func AddBook(books) {
//
//}

type Service interface {
	Create(bookData *BookData) (int64, error)
}

func Create(bookData *BookData) (int64, error) {

	//fmt.Printf("bookData------- %T \n", bookData)
	book := models.Book{}
	book.Name = bookData.Name
	book.Comment = bookData.Comment

	database.SQLDB = database.GetDB()
	result := database.SQLDB.Create(&book)
	return result.RowsAffected, result.Error
}

//func Create(bookData []*BookData) (int64, error) {
//
//	fmt.Printf("bookData------- %T \n", bookData)
//
//	//database.SQLDB = database.GetDB()
//	//result := database.SQLDB.Create(&book)
//	//return result.RowsAffected, result.Error
//	return 0, nil
//}
