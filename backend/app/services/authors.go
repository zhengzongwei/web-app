// Package          services
// @Title           authors.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/11/17 14:57

package services

import (
	"backend/app/dao"
	"backend/app/models"
	"gorm.io/gorm"
)

type DBAuthorService struct {
	DB        *gorm.DB
	AuthorDAO *dao.AuthorDAO
}

func NewAuthorService(db *gorm.DB) *DBAuthorService {
	return &DBAuthorService{
		//BookDAO:   &dao.BookDAO{DB: db, AuthorDAO: &dao.AuthorDAO{DB: db}},
		AuthorDAO: &dao.AuthorDAO{DB: db},
		DB:        db,
	}
}

func (s *DBAuthorService) CreateAuthor(authorData []models.Author) error {
	tx := s.DB.Begin()
	if err := tx.Create(authorData).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (s *DBAuthorService) ListAuthor() ([]models.Author, error) {
	return s.ListAuthor()
}

func (s *DBAuthorService) DeleteAuthor(authorIDs []uint) error {
	return s.DeleteAuthor(authorIDs)
}

func (s *DBAuthorService) EditAuthor(authorID uint, updateAuthor *models.Author) error {
	return s.EditAuthor(authorID, updateAuthor)
}

func (s *DBAuthorService) DetailAuthor(authodID uint) (*models.Author, error) {
	return s.DetailAuthor(authodID)
}
