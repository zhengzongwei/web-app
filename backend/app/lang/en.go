// Package          lang
// @Title           en.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/11/14 17:44

package lang

import (
	"backend/app/global/status_code"
)

var enText = map[int]string{
	status_code.ServerError: "Internal server error",
	status_code.OK:          "OK",
}
