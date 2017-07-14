package service

import "github.com/vetcher/comments-msv/models"

type ResponseGetCommentByID struct {
	Comment *models.Comment `json:"data"`
	Err     error           `json:"error,omitempty"`
}

type ResponsePostComment struct {
	Comment *models.Comment `json:"data"`
	Err     error           `json:"error,omitempty"`
}

type ResponseDeleteCommentByID struct {
	OK  bool  `json:"data"`
	Err error `json:"error,omitempty"`
}

type ResponseGetCommentsByAuthorID struct {
	Comments []*models.Comment `json:"data"`
	Err      error             `json:"error,omitempty"`
}

type ResponseJSON struct {
	Data interface{} `json:"data"`
	Err  string      `json:"error,omitempty"`
}
