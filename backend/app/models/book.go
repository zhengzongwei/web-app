// Package          models
// @Title           book.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/11/15 19:56

package models

import (
	"gorm.io/gorm"
)

// Book 是数据库中 "books" 表的模型
type Book struct {
	gorm.Model
	Name        string    `gorm:"column:name;not null" json:"name"`
	Comment     string    `gorm:"column:comment;type:text" json:"comment"`
	PublishDate string    `gorm:"column:publish_date;type:date;null;default:null" json:"publish_date"`
	Pages       int       `gorm:"default:0" json:"pages"`
	ISBN        string    `gorm:"unique;null" json:"isbn"`
	CoverImage  string    `gorm:"type:text" json:"cover_image"`
	Language    string    `gorm:"null" json:"language"`
	IsDelete    bool      `gorm:"default:0" json:"is_delete"`
	Authors     []Author  `gorm:"many2many:book_authors" json:"authors"`
	Publishes   []Publish `gorm:"many2many:book_publishs" json:"publishes"`
	//Publisher string `gorm:"default:null"`
}

// Author 作者表
type Author struct {
	gorm.Model
	Name  string `gorm:"column:name;not null" json:"name"`
	Phone string `gorm:"null" json:"phone"`
	Addr  string `gorm:"null" json:"addr"`
	Books []Book `gorm:"many2many:book_authors" json:"books"`
}

type Publish struct {
	gorm.Model
	Name  string `gorm:"column:name;not null" json:"name"`
	Books []Book `gorm:"many2many:book_authors" json:"books"`
}

//// Category 是数据库中 "categories" 表的模型
//type Category struct {
//	gorm.Model
//	Name        string `gorm:"unique;not null"`
//	Description string `gorm:"type:text"`
//}
//
//// BookCategory 是数据库中 "book_categories" 表的模型，用于表示图书和分类的关联关系
//type BookCategory struct {
//	gorm.Model
//	BookID     uint
//	CategoryID uint
//}
//
//// Review 是数据库中 "reviews" 表的模型，用于表示图书的评论
//type Review struct {
//	gorm.Model
//	BookID  uint
//	UserID  uint
//	Comment string `gorm:"type:text"`
//	Rating  int    `gorm:"not null"`
//}
//
//// User 是数据库中 "users" 表的模型，用于表示图书的评论者
//type User struct {
//	gorm.Model
//	Username string `gorm:"unique;not null"`
//	Email    string `gorm:"unique;not null"`
//	Password string `gorm:"not null"`
//}
//
//// Role 是数据库中 "roles" 表的模型，用于表示用户的角色
//type Role struct {
//	gorm.Model
//	Name string `gorm:"unique;not null"`
//}
//
//// UserRole 是数据库中 "user_roles" 表的模型，用于表示用户和角色的关联关系
//type UserRole struct {
//	gorm.Model
//	UserID uint
//	RoleID uint
//}
