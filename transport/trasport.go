package transport

import (
	"golang.org/x/net/context"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/vetcher/comments-msv/service/pb"
)

type GRPCServer struct {
	GetCommentByIDHandler        grpctransport.Handler
	PostCommentHandler           grpctransport.Handler
	GetCommentsByAuthorIDHandler grpctransport.Handler
	DeleteCommentByIDHandler     grpctransport.Handler
}

func (s GRPCServer) GetCommentByID(ctx context.Context, req *pb.RequestWithID) (*pb.ResponseComment, error) {
	_, rep, err := s.GetCommentByIDHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ResponseComment), nil
}

func (s GRPCServer) PostComment(ctx context.Context, req *pb.Comment) (*pb.ResponseComment, error) {
	_, rep, err := s.GetCommentByIDHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ResponseComment), nil
}

func (s GRPCServer) DeleteCommentByID(ctx context.Context, req *pb.RequestWithID) (*pb.ResponseWithBool, error) {
	_, rep, err := s.GetCommentByIDHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ResponseWithBool), nil
}

func (s GRPCServer) GetCommentsByAuthorID(ctx context.Context, req *pb.RequestWithID) (*pb.ResponseCommentsByAuthorID, error) {
	_, rep, err := s.GetCommentByIDHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ResponseCommentsByAuthorID), nil
}
