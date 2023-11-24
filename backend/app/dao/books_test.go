// Package          dao
// @Title           books_test.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/11/23 17:08

package dao

import (
	"backend/app/database"
	"backend/app/models"
	"fmt"
	"testing"
)

func TestBookDAO_BookList(t *testing.T) {
	db := database.GetDB()
	bookDao := &BookDAO{DB: db}
	list, err := bookDao.ListBook()
	if err != nil {
		return
	}
	fmt.Printf("%v", list)
}

func TestBookDAO_BookCreate(t *testing.T) {
	db := database.GetDB()
	bookDao := &BookDAO{
		DB:        db,
		AuthorDAO: &AuthorDAO{DB: db},
	}
	books := []models.Book{
		{
			Name: "Java 测试1",
			Authors: []*models.Author{
				{
					Name: "test1",
				},
				{
					Name: "test222",
				},
			},
		},
	}
	err := bookDao.CreateBook(books)
	if err != nil {
		t.Fatalf("Failed to create books: %v", err)
	}
}

func TestBookDAO_BatchBookDelete(t *testing.T) {
	var bookIDs []uint
	db := database.GetDB()
	bookDao := &BookDAO{DB: db}
	books, err := bookDao.ListBook()
	if err != nil {
		return
	}

	for _, book := range books {
		bookIDs = append(bookIDs, book.ID)
	}

	err = bookDao.DeleteBook(bookIDs)
	if err != nil {
		return
	}
}
