// Package          services
// @Title           books.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/11/16 15:49

package services

import (
	"backend/app/database"
	"backend/app/models"
	"log"
)

type BookService interface {
	BookCreate(bookData *models.Book) (int64, error)
	BookList()
}

func BookCreate(bookData *models.Book) (int64, error) {

	database.SQLDB = database.GetDB()
	result := database.SQLDB.Create(bookData)
	return int64(bookData.ID), result.Error
}

func BookList() []models.Book {
	var books []models.Book
	database.SQLDB = database.GetDB()
	result := database.SQLDB.Find(&books)
	if result.Error != nil {
		log.Printf("查询失败！%s\n", result.Error)

	}
	return books
}
