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

type commentsService struct {
	db *models.Database
}

func NewCommentService() *commentsService {
	return commentsService{}.Init()
}

func (svc *commentsService) Init() *commentsService {
	svc.db = models.InitDB()
	return svc
}

func (svc commentsService) GetCommentByID(id uint) (*models.Comment, error) {
	return svc.db.SelectCommentByID(id)
}

func (svc commentsService) PostComment(authorId uint, text string) (*models.Comment, error) {
	return svc.db.PostComment(&models.Comment{
		Text:     text,
		AuthorID: authorId,
	})
}

func (svc commentsService) GetCommentsByAuthorID(id uint) ([]*models.Comment, error) {
	return svc.db.LoadCommentsForUser(id)
}

func (svc commentsService) DeleteCommentByID(id uint) (bool, error) {
	return svc.db.DeleteComment(id)
}
