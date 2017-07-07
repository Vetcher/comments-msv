package service

import (
	"github.com/vetcher/comments-msv/models"
)

type CommentsServiceInterface interface {
	GetCommentByID(id uint) (*models.Comment, error)
	GetCommentsByAuthorId(id uint) ([]*models.Comment, error)
	PostComment(authorId uint, text string) (*models.Comment, error)
	DeleteComment(id uint) (bool, error)
}

type CommentsService struct {
	db *models.Database
}

func (svc *CommentsService) Init() {
	svc.db = models.InitDB()
}

func (svc CommentsService) GetCommentByID(id uint) (*models.Comment, error) {
	return svc.db.SelectCommentByID(id)
}

func (svc CommentsService) PostComment(authorId uint, text string) (*models.Comment, error) {
	return svc.db.PostComment(&models.Comment{
		Text:     text,
		AuthorID: authorId,
	})
}

func (svc CommentsService) GetCommentsByAuthorId(id uint) ([]*models.Comment, error) {
	return svc.db.LoadCommentsForUser(id)
}

func (svc CommentsService) DeleteComment(id uint) (bool, error) {
	return svc.db.DeleteComment(id)
}
