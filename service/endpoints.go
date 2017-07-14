package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func GetCommentEndpoint(svc CommentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RequestGetCommentByID)
		d, err := svc.GetCommentByID(req.ID)
		return &ResponseGetCommentByID{
			Comment: d,
			Err:     err,
		}, nil
	}
}

func PostCommentEndpoint(svc CommentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RequestPostComment)
		d, err := svc.PostComment(req.AuthorID, req.Text)
		return &ResponsePostComment{
			Comment: d,
			Err:     err,
		}, nil
	}
}

func DeleteCommentEndpoint(svc CommentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RequestDeleteCommentByID)
		d, err := svc.DeleteCommentByID(req.ID)
		return &ResponseDeleteCommentByID{
			OK:  d,
			Err: err,
		}, nil
	}
}

func GetCommentsByAuthorIDEndpoint(svc CommentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RequestGetCommentByAuthorID)
		d, err := svc.GetCommentsByAuthorID(req.ID)
		return &ResponseGetCommentsByAuthorID{
			Comments: d,
			Err:      err,
		}, nil
	}
}
