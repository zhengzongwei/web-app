// Package          services
// @Title           authors.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/11/17 14:57

package services

import (
	"backend/app/models"
	"gorm.io/gorm"
)

type DBAuthorService struct {
	DB *gorm.DB
}

func (s *DBAuthorService) CreateAuthor(authorData *[]models.Author) error {
	tx := s.DB.Begin()
	if err := tx.Create(*authorData).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
