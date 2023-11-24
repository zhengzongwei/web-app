// Package          models
// @Title           authors.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/11/17 14:50

package models

import "gorm.io/gorm"

type Author struct {
	Name  string  `gorm:"column:name;not null" json:"name" binding:"required"`
	Phone string  `gorm:"null" json:"phone"`
	Addr  string  `gorm:"null" json:"addr"`
	Books []*Book `gorm:"many2many:book_authors;" json:"books"`
	gorm.Model
}

type AuthorList struct {
	Authors []*Author `json:"authors"`
}
