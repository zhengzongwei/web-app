// Package          dao
// @Title           books.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/11/21 09:45

package dao

import (
	models "backend/app/models"
	"errors"
	"gorm.io/gorm"
	"log"
)

type BookDAO struct {
	DB        *gorm.DB
	AuthorDAO *AuthorDAO
}

func (d *BookDAO) GetOrCreateBook(authors []*models.Author, book models.Book) (*models.Book, error) {
	var existingBook models.Book
	tx := d.DB.Begin()
	if err := tx.Where("name = ?", book.Name).First(&existingBook).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			// 其他查询错误，回滚并返回错误
			tx.Rollback()
			log.Printf("Error querying book: %v", err)
			return &existingBook, err
		}

		// 书籍不存在，创建书籍
		book.Authors = authors // 设置关联的作者
		if err := tx.Create(&book).Error; err != nil {
			// 创建书籍失败，回滚并返回错误
			tx.Rollback()
			return &book, err
		}
	}
	return &book, nil
}

func (d *BookDAO) BookList() ([]models.Book, error) {
	var books []models.Book
	if err := d.DB.Preload("Authors").Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (d *BookDAO) BookCreate(books []models.Book) error {
	for i := range books {
		// 将每个书籍的ID设置为零，以确保GORM将其视为新的记录
		books[i].ID = 0

		authors := books[i].Authors
		// 查询已存在的作者
		var err error
		for j := range authors {
			var existingAuthor *models.Author
			if existingAuthor, err = d.AuthorDAO.GetAuthorByName(authors[j].Name); err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					return err
				}

				// 如果作者不存在，创建作者
				if err := d.AuthorDAO.CreateAuthor(authors[j]); err != nil {
					return err
				}

				// 查询新创建的作者
				if existingAuthor, err = d.AuthorDAO.GetAuthorByName(authors[j].Name); err != nil {
					return err
				}
			}

			// 选择第一个作者
			books[i].Authors[j] = existingAuthor
		}

		//// 检查是否存在相同的书籍名称和作者
		//var count int64
		//if err := d.DB.Model(&models.Book{}).
		//	Joins("JOIN authors ON books.author_id = authors.id").
		//	Where("books.name = ? AND authors.id = ?", books[i].Name, books[i].Authors[0].ID).
		//	Count(&count).Error; err != nil {
		//	return err
		//}
		//if count > 0 {
		//	return fmt.Errorf("Duplicate book with the same name and author")
		//}

		// 插入书籍
		if err := d.DB.Create(&books[i]).Error; err != nil {
			return err
		}
	}
	return nil
}
