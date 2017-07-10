package service

import (
	"github.com/vetcher/comments-msv/models"
)

type CommentService interface {
	GetCommentByID(id uint) (*models.Comment, error)
	GetCommentsByAuthorID(id uint) ([]*models.Comment, error)
	PostComment(authorId uint, text string) (*models.Comment, error)
	DeleteCommentByID(id uint) (bool, error)
}

type commentService struct {
	db *models.Database
}

func NewCommentService(db *models.Database) CommentService {
	return &commentService{db: db}
}

func (svc commentService) GetCommentByID(id uint) (*models.Comment, error) {
	return svc.db.SelectCommentByID(id)
}

func (svc commentService) PostComment(authorId uint, text string) (*models.Comment, error) {
	return svc.db.PostComment(&models.Comment{
		Text:     text,
		AuthorID: authorId,
	})
}

func (svc commentService) GetCommentsByAuthorID(id uint) ([]*models.Comment, error) {
	return svc.db.LoadCommentsForUser(id)
}

func (svc commentService) DeleteCommentByID(id uint) (bool, error) {
	return svc.db.DeleteComment(id)
}
