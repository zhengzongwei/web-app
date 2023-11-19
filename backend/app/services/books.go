// Package          services
// @Title           books.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/11/16 15:49

package services

import (
	"backend/app/models"
	"errors"
	"gorm.io/gorm"
	"log"
)

type DBBookService struct {
	DB *gorm.DB
}

func (s *DBBookService) GetOrCreateAuthor(tx *gorm.DB, authorName string) (*models.Author, error) {
	var existingAuthor models.Author
	if err := tx.Where("name = ?", authorName).First(&existingAuthor).Error; err != nil {
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

	return &existingAuthor, nil
}

func (s *DBBookService) GetOrCreateBook(tx *gorm.DB, authors []*models.Author, book models.Book) (*models.Book, error) {
	var existingBook models.Book
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

func (s *DBBookService) BookCreate(bookData *[]models.Book) error {
	tx := s.DB.Begin()

	for _, book := range *bookData {
		// 处理作者
		var authors []*models.Author
		for _, author := range book.Authors {
			existingAuthor, err := s.GetOrCreateAuthor(tx, author.Name)
			if err != nil {
				return err
			}
			authors = append(authors, existingAuthor)
		}

		// 检查书籍是否已存在
		_, err := s.GetOrCreateBook(tx, authors, book)
		if err != nil {
			return err
		}
	}

	// 提交事务
	tx.Commit()
	return nil
}

func (s *DBBookService) BookList() *[]models.Book {
	var books []models.Book

	result := s.DB.Preload("Authors").Find(books)
	if result.Error != nil {
		log.Printf("查询失败！%s\n", result.Error)
	}
	return &books
}

func (s *DBBookService) BookDetail(bookId uint) models.Book {
	var book models.Book

	result := s.DB.First(&book, bookId)
	if result.Error != nil {
		log.Printf("查询失败！%s\n", result.Error)
	}
	return book
}

func (s *DBBookService) BookDelete(books *[]models.Book) (int64, error) {
	// 使用软删除
	tx := *s.DB.Begin()
	for _, book := range *books {
		// 	预加载关联的作者数据
		if err := tx.Preload("Authors").First(book, book.ID).Error; err != nil {
			tx.Rollback()
			return 0, err
		}

		if err := tx.Model(book).Association("Authors").Clear(); err != nil {
			tx.Rollback()
			return 0, err
		}

		if err := tx.Delete(book).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}
	tx.Commit()
	return int64(len(*books)), nil

}
