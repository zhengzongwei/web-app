// Package          response
// @Title           resp.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/11/15 15:44

package response

import (
	"backend/app/lang"
	"github.com/gin-gonic/gin"
)

//func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
//	ctx.JSON(httpStatus, gin.H{
//		"code": code,
//		"msg":  msg,
//		"data": data,
//	})
//}

func Response(ctx *gin.Context, httpStatus int, code int, data gin.H) {
	ctx.JSON(httpStatus, gin.H{
		"code": code,
		"msg":  lang.Text(code),
		"data": data,
	})
}
