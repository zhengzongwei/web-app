// Package          v1
// @Title           controller.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/11/15 16:20

package v1

import (
	"backend/app/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexApi(c *gin.Context) {
	response.Response(c, http.StatusOK, 0, nil)
}
