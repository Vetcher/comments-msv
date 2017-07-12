package client

import (
	"context"

	"errors"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/vetcher/comments-msv/models"
	"github.com/vetcher/comments-msv/service"
	"github.com/vetcher/comments-msv/service/pb"
	"google.golang.org/grpc"
)

type endpoints struct {
	GetCommentByIDEndpoint        endpoint.Endpoint
	PostCommentEndpoint           endpoint.Endpoint
	DeleteCommentByIDEndpoint     endpoint.Endpoint
	GetCommentsByAuthorIDEndpoint endpoint.Endpoint
}

func (e endpoints) GetCommentByID(id uint) (*models.Comment, error) {
	req := service.RequestOnlyWithID{ID: id}
	resp, err := e.GetCommentByIDEndpoint(context.Background(), req)
	if err != nil {
		return nil, err
	}
	// возможно, errors.New не нужна...
	return resp.(service.JsonResponse).Data.(*models.Comment), errors.New(resp.(service.JsonResponse).Err)
}

func (e endpoints) PostComment(authorId uint, text string) (*models.Comment, error) {
	req := service.RequestPostComment{Text: text, AuthorID: authorId}
	resp, err := e.PostCommentEndpoint(context.Background(), req)
	if err != nil {
		return nil, err
	}
	// возможно, errors.New не нужна...
	return resp.(service.JsonResponse).Data.(*models.Comment), errors.New(resp.(service.JsonResponse).Err)
}

func (e endpoints) DeleteCommentByID(id uint) (bool, error) {
	req := service.RequestOnlyWithID{ID: id}
	resp, err := e.DeleteCommentByIDEndpoint(context.Background(), req)
	if err != nil {
		return nil, err
	}
	// возможно, errors.New не нужна...
	return resp.(service.JsonResponse).Data.(bool), errors.New(resp.(service.JsonResponse).Err)
}

func (e endpoints) GetCommentsByAuthorID(id uint) ([]*models.Comment, error) {
	req := service.RequestOnlyWithID{ID: id}
	resp, err := e.GetCommentsByAuthorIDEndpoint(context.Background(), req)
	if err != nil {
		return nil, err
	}
	// возможно, errors.New не нужна...
	return resp.(service.JsonResponse).Data.([]*models.Comment), errors.New(resp.(service.JsonResponse).Err)
}

func NewClient(conn *grpc.ClientConn) service.CommentService {
	getCommentByIDEndpoint := grpctransport.NewClient(
		conn,
		"pb.CommentSVC",
		"GetCommentByID",
		service.DecodeGRPCRequestWithID,
		service.EncodeGRPCResponseComment,
		pb.ResponseComment{},
	).Endpoint()
	postCommentEndpoint := grpctransport.NewClient(
		conn,
		"pb.CommentSVC",
		"PostComment",
		service.DecodeGRPCRequestPostComment,
		service.EncodeGRPCResponseComment,
		pb.ResponseComment{},
	).Endpoint()
	deleteCommentByIDEndpoint := grpctransport.NewClient(
		conn,
		"pb.CommentSVC",
		"DeleteCommentByID",
		service.DecodeGRPCRequestWithID,
		service.EncodeGRPCResponseBool,
		pb.ResponseWithBool{},
	).Endpoint()
	getCommentsByAuthorIDEndpoint := grpctransport.NewClient(
		conn,
		"pb.CommentSVC",
		"GetCommentsByAuthorID",
		service.DecodeGRPCRequestWithID,
		service.EncodeGRPCResponseCommentsByAuthorID,
		pb.ResponseWithBool{},
	).Endpoint()

	return endpoints{
		GetCommentByIDEndpoint:        getCommentByIDEndpoint,
		PostCommentEndpoint:           postCommentEndpoint,
		DeleteCommentByIDEndpoint:     deleteCommentByIDEndpoint,
		GetCommentsByAuthorIDEndpoint: getCommentsByAuthorIDEndpoint,
	}
}
