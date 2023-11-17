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

func TestService_Create(t *testing.T) {
	book := models.Book{
		Name:    "C#程序设计",
		Comment: "测试123",
	}

	result, err := BookCreate(&book)
	if err != nil {
		return
	}
	fmt.Printf("%d \n", result)
}
