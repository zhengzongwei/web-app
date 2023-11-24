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
	list, err := bookDao.BookList()
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
			Name: "Java 测试",
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
	err := bookDao.BookCreate(books)
	if err != nil {
		t.Fatalf("Failed to create books: %v", err)
	}
}
