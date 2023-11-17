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
	BookCreate(bookData *models.Book) (uint, error)
	BookList()
	BookDetail(bookID models.Book)
}

func BookCreate(bookData *models.Book) (uint, error) {
	database.SQLDB = database.GetDB()
	result := database.SQLDB.Create(bookData)
	return bookData.ID, result.Error
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

func BookDetail(bookId uint) models.Book {
	var book models.Book
	database.SQLDB = database.GetDB()
	result := database.SQLDB.First(&book, bookId)
	if result.Error != nil {
		log.Printf("查询失败！%s\n", result.Error)
	}
	return book
}

func BookDelete(books *models.Book) (int64, error) {
	// 使用软删除
	//for _, v := range bookIds {
	//
	//}
	database.SQLDB = database.GetDB()
	result := database.SQLDB.Delete(&books)
	if result.Error != nil {
		log.Printf("删除失败！%s\n", result.Error)
	}
	// 强制删除
	return result.RowsAffected, result.Error
}
