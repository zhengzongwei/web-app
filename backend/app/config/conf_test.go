// Package          config
// @Title           conf_test.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/11/14 17:41

package config

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	config := Get()
	fmt.Print(config)
}
