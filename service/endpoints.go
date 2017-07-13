package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func GetCommentEndpoint(svc CommentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RequestGetCommentByID)
		d, err := svc.GetCommentByID(req.ID)
		if err != nil {
			return &Response{nil, err}, nil
		}
		return &Response{d, nil}, nil
	}
}

func PostCommentEndpoint(svc CommentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RequestPostComment)
		d, err := svc.PostComment(req.AuthorID, req.Text)
		if err != nil {
			return &Response{nil, err}, nil
		}
		return &Response{d, nil}, nil
	}
}

func DeleteCommentEndpoint(svc CommentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RequestDeleteCommentByID)
		d, err := svc.DeleteCommentByID(req.ID)
		if err != nil {
			return &Response{nil, err}, nil
		}
		return &Response{d, nil}, nil
	}
}

func GetCommentsByAuthorIDEndpoint(svc CommentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RequestGetCommentByAuthorID)
		d, err := svc.GetCommentsByAuthorID(req.ID)
		if err != nil {
			return &Response{nil, err}, nil
		}
		return &Response{d, nil}, nil
	}
}
