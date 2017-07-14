package transport

import (
	"context"

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
	req := service.RequestGetCommentByID{ID: id}
	resp, err := e.GetCommentByIDEndpoint(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	input := resp.(*service.ResponseGetCommentByID)
	return input.Comment, input.Err
}

func (e endpoints) PostComment(authorId uint, text string) (*models.Comment, error) {
	req := service.RequestPostComment{Text: text, AuthorID: authorId}
	resp, err := e.PostCommentEndpoint(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	input := resp.(*service.ResponsePostComment)
	return input.Comment, input.Err
}

func (e endpoints) DeleteCommentByID(id uint) (bool, error) {
	req := service.RequestDeleteCommentByID{ID: id}
	resp, err := e.DeleteCommentByIDEndpoint(context.Background(), &req)
	if err != nil {
		return false, err
	}
	return resp.(*service.ResponseDeleteCommentByID).OK, resp.(*service.ResponseDeleteCommentByID).Err
}

func (e endpoints) GetCommentsByAuthorID(id uint) ([]*models.Comment, error) {
	req := service.RequestGetCommentByAuthorID{ID: id}
	resp, err := e.GetCommentsByAuthorIDEndpoint(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	input := resp.(*service.ResponseGetCommentsByAuthorID)
	return input.Comments, input.Err
}

func NewClient(conn *grpc.ClientConn) service.CommentService {
	getCommentByIDEndpoint := grpctransport.NewClient(
		conn,
		"pb.CommentSVC",
		"GetCommentByID",
		service.EncodeGRPCRequestGetCommentByID,
		service.DecodeGRPCResponseGetCommentByID,
		pb.ResponseComment{},
	).Endpoint()
	postCommentEndpoint := grpctransport.NewClient(
		conn,
		"pb.CommentSVC",
		"PostComment",
		service.EncodeGRPCRequestPostComment,
		service.DecodeGRPCResponsePostComment,
		pb.ResponseComment{},
	).Endpoint()
	deleteCommentByIDEndpoint := grpctransport.NewClient(
		conn,
		"pb.CommentSVC",
		"DeleteCommentByID",
		service.EncodeGRPCRequestDeleteCommentByID,
		service.DecodeGRPCResponseDeleteCommentByID,
		pb.ResponseDeleteByID{},
	).Endpoint()
	getCommentsByAuthorIDEndpoint := grpctransport.NewClient(
		conn,
		"pb.CommentSVC",
		"GetCommentsByAuthorID",
		service.EncodeGRPCRequestGetCommentByAuthorID,
		service.DecodeGRPCResponseCommentsByAuthorID,
		pb.ResponseCommentsByAuthorID{},
	).Endpoint()

	return endpoints{
		GetCommentByIDEndpoint:        getCommentByIDEndpoint,
		PostCommentEndpoint:           postCommentEndpoint,
		DeleteCommentByIDEndpoint:     deleteCommentByIDEndpoint,
		GetCommentsByAuthorIDEndpoint: getCommentsByAuthorIDEndpoint,
	}
}
