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
		Name:    "C#程序设计",
		Comment: "测试123",
	}
	//book := []*BookData{
	//	{
	//		Name:    "C#程序设计",
	//		Comment: "测试123",
	//	},
	//	{
	//		Name:    "C#程序设计111",
	//		Comment: "测试123",
	//	},
	//}
	result, err := Create(&book)
	if err != nil {
		return
	}
	fmt.Printf("%d \n", result)
}
