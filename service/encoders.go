package service

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/vetcher/comments-msv/models"
	"github.com/vetcher/comments-msv/service/pb"
	"github.com/vetcher/comments-msv/util"
)

func ServeJSON(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func EncodeGRPCResponseComment(_ context.Context, svcResp interface{}) (interface{}, error) {
	resp := svcResp.(*Response)
	if resp.Data == nil {
		return &pb.ResponseComment{
			Data: nil,
			Err:  util.Err2Str(resp.Err),
		}, nil
	}
	comment := resp.Data.(*models.Comment)
	return &pb.ResponseComment{
		Data: &pb.Comment{
			Id:        uint32(comment.ID),
			Text:      comment.Text,
			AuthorId:  uint32(comment.AuthorID),
			CreatedAt: comment.CreatedAt.Unix(),
		},
		Err: util.Err2Str(resp.Err),
	}, nil
}

func EncodeGRPCResponseDeleteCommentByID(_ context.Context, svcResp interface{}) (interface{}, error) {
	resp := svcResp.(*Response)
	ok := resp.Data.(bool)
	return &pb.ResponseDeleteByID{
		Ok:  ok,
		Err: util.Err2Str(resp.Err),
	}, nil
}

func EncodeGRPCResponseCommentsByAuthorID(_ context.Context, svcResp interface{}) (interface{}, error) {
	resp := svcResp.(*Response)
	return &pb.ResponseCommentsByAuthorID{
		Comments: util.ConvDBCommentsToPBComments(resp.Data.([]*models.Comment)),
		Err:      util.Err2Str(resp.Err),
	}, nil
}

func EncodeGRPCRequestGetCommentByID(_ context.Context, svcReq interface{}) (interface{}, error) {
	req := svcReq.(*RequestGetCommentByID)
	return &pb.RequestGetCommentByID{Id: uint32(req.ID)}, nil
}

func EncodeGRPCRequestGetCommentByAuthorID(_ context.Context, svcReq interface{}) (interface{}, error) {
	req := svcReq.(*RequestGetCommentByAuthorID)
	return &pb.RequestGetCommentByAuthorID{Id: uint32(req.ID)}, nil
}

func EncodeGRPCRequestDeleteCommentByID(_ context.Context, svcReq interface{}) (interface{}, error) {
	req := svcReq.(*RequestDeleteCommentByID)
	return &pb.RequestDeleteCommentByID{Id: uint32(req.ID)}, nil
}

func EncodeGRPCRequestPostComment(_ context.Context, svcReq interface{}) (interface{}, error) {
	req := svcReq.(*RequestPostComment)
	return &pb.Comment{AuthorId: uint32(req.AuthorID), Text: req.Text}, nil
}
