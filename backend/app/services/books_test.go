// Package          services
// @Title           books_test.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/11/16 16:56

package services

import (
	"fmt"
	"testing"
)

func TestService_Create(t *testing.T) {
	book := BookData{
		Name:    "C#程dd序设计",
		Comment: "test1ss23",
	}

	result, err := Create(&book)
	if err != nil {
		return
	}
	fmt.Printf("%d \n", result)
}
