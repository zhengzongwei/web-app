// Package          env
// @Title           env.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/11/14 17:26

package env

import (
	"fmt"
	"strings"
)

var (
	active Environment
	dev    Environment = &environment{value: "dev"}
	//fat    Environment = &environment{value: "fat"}
	//uat    Environment = &environment{value: "uat"}
	prod Environment = &environment{value: "prod"}
)

var _ Environment = (*environment)(nil)

type Environment interface {
	Value() string
	IsDev() bool
	//IsFat() bool
	//IsUat() bool
	IsProd() bool
	t()
}

type environment struct {
	value string
}

func (e *environment) Value() string {
	return e.value
}

func (e *environment) IsDev() bool {
	return e.value == "dev"
}

//func (e *environment) IsFat() bool {
//	return e.value == "fat"
//}
//
//func (e *environment) IsUat() bool {
//	return e.value == "uat"
//}

func (e *environment) IsProd() bool {
	return e.value == "prod"
}

func (e *environment) t() {}

func init() {
	//env := flag.String("env", "", "请输入运行环境:\n dev:开发环境\n fat:测试环境\n uat:预上线环境\n prod:正式环境\n")
	//flag.Parse()

	var env string = "dev"
	//env := "dev"
	//switch strings.ToLower(strings.TrimSpace(*env)) {
	switch strings.ToLower(strings.TrimSpace(env)) {
	case "dev":
		active = dev
	//case "fat":
	//	active = fat
	//case "uat":
	//	active = uat
	case "prod":
		active = prod
	default:
		active = dev
		fmt.Println("Warning: '-env' cannot be found, or it is illegal. The default 'fat' will be used.")
	}
}

// Active 当前配置的env
func Active() Environment {
	return active
}
