// Package          models
// @Title           publish.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/11/17 14:50

package models

import "gorm.io/gorm"

type Publish struct {
	gorm.Model
	Name  string `gorm:"column:name;not null" json:"name"`
	Books []Book `gorm:"many2many:book_authors" json:"books"`
}
