package service

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/vetcher/comments-msv/service/pb"
	"github.com/vetcher/comments-msv/util"
)

func ServeJSON(_ context.Context, w http.ResponseWriter, response interface{}) error {
	switch response.(type) {
	case *ResponseGetCommentByID:
		response = ResponseJSON{
			Data: response.(*ResponseGetCommentByID).Comment,
			Err:  util.Err2Str(response.(*ResponseGetCommentByID).Err),
		}
	case *ResponsePostComment:
		response = ResponseJSON{
			Data: response.(*ResponsePostComment).Comment,
			Err:  util.Err2Str(response.(*ResponsePostComment).Err),
		}
	case *ResponseGetCommentsByAuthorID:
		response = ResponseJSON{
			Data: response.(*ResponseGetCommentsByAuthorID).Comments,
			Err:  util.Err2Str(response.(*ResponseGetCommentsByAuthorID).Err),
		}
	case *ResponseDeleteCommentByID:
		response = ResponseJSON{
			Data: response.(*ResponseDeleteCommentByID).OK,
			Err:  util.Err2Str(response.(*ResponseDeleteCommentByID).Err),
		}
	}
	return json.NewEncoder(w).Encode(response)
}

func EncodeGRPCResponseGetCommentByID(_ context.Context, svcResp interface{}) (interface{}, error) {
	resp := svcResp.(*ResponseGetCommentByID)
	if resp.Comment == nil {
		return &pb.ResponseComment{
			Data: nil,
			Err:  util.Err2Str(resp.Err),
		}, nil
	}
	return &pb.ResponseComment{
		Data: &pb.Comment{
			Id:        uint32(resp.Comment.ID),
			Text:      resp.Comment.Text,
			AuthorId:  uint32(resp.Comment.AuthorID),
			CreatedAt: resp.Comment.CreatedAt.Unix(),
		},
		Err: util.Err2Str(resp.Err),
	}, nil
}

func EncodeGRPCResponsePostComment(_ context.Context, svcResp interface{}) (interface{}, error) {
	resp := svcResp.(*ResponsePostComment)
	if resp.Comment == nil {
		return &pb.ResponseComment{
			Data: nil,
			Err:  util.Err2Str(resp.Err),
		}, nil
	}
	return &pb.ResponseComment{
		Data: &pb.Comment{
			Id:        uint32(resp.Comment.ID),
			Text:      resp.Comment.Text,
			AuthorId:  uint32(resp.Comment.AuthorID),
			CreatedAt: resp.Comment.CreatedAt.Unix(),
		},
		Err: util.Err2Str(resp.Err),
	}, nil
}

func EncodeGRPCResponseDeleteCommentByID(_ context.Context, svcResp interface{}) (interface{}, error) {
	resp := svcResp.(*ResponseDeleteCommentByID)
	return &pb.ResponseDeleteByID{
		Ok:  resp.OK,
		Err: util.Err2Str(resp.Err),
	}, nil
}

func EncodeGRPCResponseCommentsByAuthorID(_ context.Context, svcResp interface{}) (interface{}, error) {
	resp := svcResp.(*ResponseGetCommentsByAuthorID)
	return &pb.ResponseCommentsByAuthorID{
		Comments: util.ConvertDatabase2PBComments(resp.Comments),
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
