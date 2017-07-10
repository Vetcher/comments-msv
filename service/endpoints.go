package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type JsonResponse struct {
	Data interface{} `json:"data"`
	Err  string      `json:"error,omitempty"`
}

func GetCommentEndpoint(svc CommentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RequestOnlyWithId)
		d, err := svc.GetCommentByID(req.ID)
		if err != nil {
			return &JsonResponse{nil, err.Error()}, nil
		}
		return &JsonResponse{d, ""}, nil
	}
}

func PostCommentEndpoint(svc CommentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RequestPostComment)
		d, err := svc.PostComment(req.AuthorID, req.Text)
		if err != nil {
			return &JsonResponse{nil, err.Error()}, nil
		}
		return &JsonResponse{d, ""}, nil
	}
}

func DeleteCommentEndpoint(svc CommentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RequestOnlyWithId)
		d, err := svc.DeleteCommentByID(req.ID)
		if err != nil {
			return &JsonResponse{nil, err.Error()}, nil
		}
		return &JsonResponse{d, ""}, nil
	}
}

func GetCommentsByAuthorIDEndpoint(svc CommentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RequestOnlyWithId)
		d, err := svc.GetCommentsByAuthorID(req.ID)
		if err != nil {
			return &JsonResponse{nil, err.Error()}, nil
		}
		return &JsonResponse{d, ""}, nil
	}
}
