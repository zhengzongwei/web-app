// Package          dao
// @Title           authors.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/11/21 09:45

package dao

import (
	"backend/app/models"
	"errors"
	"gorm.io/gorm"
	"log"
)

type AuthorDAO struct {
	DB *gorm.DB
}

func (d *AuthorDAO) GetOrCreateAuthor(authorName string) (*models.Author, error) {
	var existingAuthor models.Author
	tx := d.DB.Begin()
	if err := tx.Where("name = ?", authorName).Find(&existingAuthor).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			// 其他查询错误，回滚并返回错误
			tx.Rollback()
			log.Printf("Error querying author: %v", err)
			return nil, err
		}

		// 作者不存在，创建作者
		if err := tx.Where(models.Author{Name: authorName}).FirstOrCreate(&existingAuthor).Error; err != nil {
			// 创建作者失败，回滚并返回错误
			tx.Rollback()
			return nil, err
		}
	}
	tx.Commit()
	return &existingAuthor, nil
}

func (d *AuthorDAO) GetAuthorByName(name string) (*models.Author, error) {
	var author models.Author
	err := d.DB.Preload("Books").Where("name = ?", name).First(&author).Error
	// 判断作者是否存在
	if author.ID == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &author, err
}

func (d *AuthorDAO) ListAuthor() ([]models.Author, error) {
	var authors []models.Author
	if err := d.DB.Find(&authors).Error; err != nil {
		return nil, err
	}
	return authors, nil

}

func (d *AuthorDAO) CreateAuthor(author *models.Author) error {
	return d.DB.Create(author).Error
}

func (d *AuthorDAO) EditAuthor(authorID uint, updateAuthor *models.Author) error {
	var author models.Author
	tx := d.DB.Begin()

	if err := tx.First(&author, authorID).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 更新作者信息
	if updateAuthor.Name != "" {
		author.Name = updateAuthor.Name
	}

	if updateAuthor.Phone != "" {
		author.Phone = updateAuthor.Phone
	}

	if updateAuthor.Addr != "" {
		author.Addr = updateAuthor.Addr
	}
	//
	//if updateAuthor.Books != ""{
	//
	//}

	if err := tx.Save(&author).Commit().Error; err != nil {
		return err
	}
	return nil
}

func (d *AuthorDAO) DeleteAuthor(authorIDs []uint) error {
	// 开启事务
	tx := d.DB.Begin()

	for _, authorID := range authorIDs {
		// 检查作者是否有关联的书籍
		var bookCount int64
		if err := tx.Model(&models.Book{}).
			Joins("JOIN book_authors ON books.id = book_authors.book_id").
			Where("book_authors.author_id = ?", authorID).
			Count(&bookCount).Error; err != nil {
			tx.Rollback()
			return err
		}

		if bookCount == 0 {
			// 如果作者没有关联的书籍，删除作者记录
			if err := tx.Delete(&models.Author{}, authorID).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	// 提交事务
	tx.Commit()

	return nil
}

func (d *AuthorDAO) GetAuthorById(id uint) (*models.Author, error) {
	var author models.Author
	err := d.DB.Preload("Books").Where("id = ?", id).First(&author).Error
	return &author, err
}
