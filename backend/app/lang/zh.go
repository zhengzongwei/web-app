// Package          lang
// @Title           zh.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/11/14 17:44

package lang

import (
	"backend/app/global/status_code"
)

var zhCNText = map[int]string{
	status_code.ServerError: "内部服务器错误",
	status_code.OK:          "请求成功",
}
