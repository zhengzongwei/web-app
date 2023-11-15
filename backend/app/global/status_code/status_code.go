// Package          app
// @Title           status_code.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/11/15 15:37

package status_code

type Failure struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

const (
	ServerError = -1
	OK          = 0
)
