// Package          services
// @Title           books.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/11/16 15:49

package services

import (
	"backend/app/database"
	"backend/app/models"
)

type BookService interface {
	BookCreate(bookData *models.Book) (int64, error)
}

func BookCreate(bookData *models.Book) (uint, error) {

	database.SQLDB = database.GetDB()
	result := database.SQLDB.Create(bookData)
	return bookData.ID, result.Error
}
