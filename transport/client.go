package transport

import (
	"context"

	"time"

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
	input := resp.(service.Response).Data.(*pb.Comment)
	if input == nil {
		return nil, resp.(service.Response).Err
	}
	return &models.Comment{
		Text:      input.Text,
		AuthorID:  uint(input.AuthorId),
		ID:        uint(input.Id),
		CreatedAt: time.Unix(input.CreatedAt, 0),
	}, resp.(service.Response).Err
}

func (e endpoints) PostComment(authorId uint, text string) (*models.Comment, error) {
	req := service.RequestPostComment{Text: text, AuthorID: authorId}
	resp, err := e.PostCommentEndpoint(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	data := resp.(service.Response).Data.(*pb.Comment)
	if data == nil {
		return nil, resp.(service.Response).Err
	}
	return &models.Comment{
		Text:      data.Text,
		AuthorID:  uint(data.AuthorId),
		ID:        uint(data.Id),
		CreatedAt: time.Unix(data.CreatedAt, 0),
	}, resp.(service.Response).Err
}

func (e endpoints) DeleteCommentByID(id uint) (bool, error) {
	req := service.RequestDeleteCommentByID{ID: id}
	resp, err := e.DeleteCommentByIDEndpoint(context.Background(), &req)
	if err != nil {
		return false, err
	}
	return resp.(service.Response).Data.(bool), resp.(service.Response).Err
}

func (e endpoints) GetCommentsByAuthorID(id uint) ([]*models.Comment, error) {
	req := service.RequestGetCommentByAuthorID{ID: id}
	resp, err := e.GetCommentsByAuthorIDEndpoint(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	data := resp.(service.Response).Data.([]*pb.Comment)
	return ConvPBToModelComments(data), resp.(service.Response).Err
}

func ConvPBToModelComments(data []*pb.Comment) []*models.Comment {
	var comments []*models.Comment
	for _, x := range data {
		comments = append(comments, &models.Comment{
			AuthorID:  uint(x.AuthorId),
			Text:      x.Text,
			ID:        uint(x.Id),
			CreatedAt: time.Unix(x.CreatedAt, 0),
		})
	}
	return comments
}

func NewClient(conn *grpc.ClientConn) service.CommentService {
	getCommentByIDEndpoint := grpctransport.NewClient(
		conn,
		"pb.CommentSVC",
		"GetCommentByID",
		service.EncodeGRPCRequestGetCommentByID,
		service.DecodeGRPCResponseComment,
		pb.ResponseComment{},
	).Endpoint()
	postCommentEndpoint := grpctransport.NewClient(
		conn,
		"pb.CommentSVC",
		"PostComment",
		service.EncodeGRPCRequestPostComment,
		service.DecodeGRPCResponseComment,
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
