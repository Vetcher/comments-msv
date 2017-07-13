package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func GetCommentEndpoint(svc CommentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RequestOnlyWithID)
		d, err := svc.GetCommentByID(req.ID)
		if err != nil {
			return &Response{nil, err.Error()}, nil
		}
		return &Response{d, ""}, nil
	}
}

func PostCommentEndpoint(svc CommentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RequestPostComment)
		d, err := svc.PostComment(req.AuthorID, req.Text)
		if err != nil {
			return &Response{nil, err.Error()}, nil
		}
		return &Response{d, ""}, nil
	}
}

func DeleteCommentEndpoint(svc CommentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RequestOnlyWithID)
		d, err := svc.DeleteCommentByID(req.ID)
		if err != nil {
			return &Response{nil, err.Error()}, nil
		}
		return &Response{d, ""}, nil
	}
}

func GetCommentsByAuthorIDEndpoint(svc CommentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RequestOnlyWithID)
		d, err := svc.GetCommentsByAuthorID(req.ID)
		if err != nil {
			return &Response{nil, err.Error()}, nil
		}
		return &Response{d, ""}, nil
	}
}
