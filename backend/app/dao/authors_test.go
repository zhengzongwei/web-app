// Package          dao
// @Title           authors_test.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/11/23 17:00

package dao

import (
	"backend/app/database"
	"backend/app/models"
	"fmt"
	"testing"
)

func TestAuthorDAO_GetAuthorByName(t *testing.T) {
	db := database.GetDB()
	authorDAO := AuthorDAO{DB: db}
	author, err := authorDAO.GetAuthorByName("test1")
	if err != nil {
		return
	}
	fmt.Printf("%v\n", *author)
}
func TestAuthorDAO_AuthorList(t *testing.T) {
	db := database.GetDB()
	authorDAO := AuthorDAO{DB: db}
	authors, err := authorDAO.ListAuthor()
	if err != nil {
		return
	}
	fmt.Printf("author info %v", authors)
}

func TestAuthorDAO_EditAuthor(t *testing.T) {
	db := database.GetDB()
	authorDAO := AuthorDAO{DB: db}
	var authorID uint = 137
	authorData := models.Author{
		Name: "徐波",
	}
	if err := authorDAO.EditAuthor(authorID, &authorData).Error; err != nil {

	}

}

func TestAuthorDAO_DeleteAuthor(t *testing.T) {
	db := database.GetDB()
	authorDAO := AuthorDAO{DB: db}
	authorIDs := []uint{
		131,
		132,
		133,
		134,
	}
	err := authorDAO.DeleteAuthor(authorIDs)
	if err != nil {
		return
	}

}
