package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type JsonResponse struct {
	Data interface{} `json:"data"`
	//Err  string      `json:"error,omitempty"`
}

func GetCommentEndpoint(svc CommentsServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RequestOnlyWithId)
		d, err := svc.GetCommentByID(req.ID)
		if err != nil {
			return nil, err
		}
		return JsonResponse{d}, nil
	}
}

func PostCommentEndpoint(svc CommentsServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RequestPostComment)
		d, err := svc.PostComment(req.AuthorID, req.Text)
		if err != nil {
			return nil, err
		}
		return JsonResponse{d}, nil
	}
}

func DeleteCommentEndpoint(svc CommentsServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RequestOnlyWithId)
		d, err := svc.DeleteComment(req.ID)
		if err != nil {
			return nil, err
		}
		return JsonResponse{d}, nil
	}
}

func GetCommentsForUserEndpoint(svc CommentsServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RequestOnlyWithId)
		d, err := svc.GetCommentsByAuthorId(req.ID)
		if err != nil {
			return nil, err
		}
		return JsonResponse{d}, nil
	}
}
