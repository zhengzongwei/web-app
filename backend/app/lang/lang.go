// Package          lang
// @Title           lang.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/11/15 15:41

package lang

import (
	"backend/app/config"
)

func Text(code int) string {
	lang := config.Get().Language.Local

	switch lang {
	case config.EnUS:
		return enText[code]
	case config.ZhCN:
		return zhCNText[code]
	default:
		return zhCNText[code]
	}
}
