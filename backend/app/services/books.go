// Package          services
// @Title           books.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/11/16 15:49

package services

import (
	dao "backend/app/dao"
	models "backend/app/models"
	"gorm.io/gorm"
)

type BookService struct {
	BookDAO   *dao.BookDAO
	AuthorDAO *dao.AuthorDAO
	DB        *gorm.DB
}

func NewBookService(db *gorm.DB) *BookService {
	return &BookService{
		BookDAO:   &dao.BookDAO{DB: db, AuthorDAO: &dao.AuthorDAO{DB: db}},
		AuthorDAO: &dao.AuthorDAO{DB: db},
		DB:        db,
	}
}

func (s *BookService) GetOrCreateAuthor(authorName string) (*models.Author, error) {
	return s.AuthorDAO.GetOrCreateAuthor(authorName)
}

func (s *BookService) GetOrCreateBook(authors []*models.Author, book models.Book) (*models.Book, error) {
	return s.BookDAO.GetOrCreateBook(authors, book)
}

func (s *BookService) CreateBook(bookData []models.Book) error {
	return s.BookDAO.CreateBook(bookData)
}

func (s *BookService) ListBook() ([]models.Book, error) {
	return s.BookDAO.ListBook()
}

func (s *BookService) DeleteBook(bookIDs []uint) error {
	return s.BookDAO.DeleteBook(bookIDs)
}

func (s *BookService) DetailBook(bookID uint) (*models.Book, error) {
	return s.BookDAO.GetBookById(bookID)
}

func (s *BookService) EditBook(bookID uint, updatedBook *models.Book) error {
	return s.BookDAO.EditBook(bookID, updatedBook)
}
