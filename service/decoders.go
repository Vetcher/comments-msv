package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vetcher/comments-msv/service/pb"
	"github.com/vetcher/comments-msv/util"
)

func DecodeRequestGetCommentByID(_ context.Context, r *http.Request) (interface{}, error) {
	var request RequestGetCommentByID
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, fmt.Errorf("parsing error: %v", err)
	}
	return request, nil
}

func DecodeRequestGetCommentByAuthorID(_ context.Context, r *http.Request) (interface{}, error) {
	var request RequestGetCommentByAuthorID
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, fmt.Errorf("parsing error: %v", err)
	}
	return request, nil
}

func DecodeRequestDeleteCommentByID(_ context.Context, r *http.Request) (interface{}, error) {
	var request RequestDeleteCommentByID
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, fmt.Errorf("parsing error: %v", err)
	}
	return request, nil
}

func DecodeRequestPostComment(_ context.Context, r *http.Request) (interface{}, error) {
	var request RequestPostComment
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, fmt.Errorf("parsing error: %v", err)
	}
	return request, nil
}

func DecodeGRPCRequestGetCommentByID(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.RequestGetCommentByID)
	return RequestGetCommentByID{
		ID: uint(req.Id),
	}, nil
}
func DecodeGRPCRequestDeleteCommentByID(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.RequestDeleteCommentByID)
	return RequestDeleteCommentByID{
		ID: uint(req.Id),
	}, nil
}
func DecodeGRPCRequestGetCommentByAuthorID(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.RequestGetCommentByAuthorID)
	return RequestGetCommentByAuthorID{
		ID: uint(req.Id),
	}, nil
}

func DecodeGRPCRequestPostComment(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.Comment)
	return RequestPostComment{
		Text:     req.Text,
		AuthorID: uint(req.AuthorId),
	}, nil
}

func DecodeGRPCResponseComment(_ context.Context, grpcResp interface{}) (interface{}, error) {
	resp := grpcResp.(*pb.ResponseComment)
	return Response{
		Data: resp.Data,
		Err:  util.Str2Err(resp.Err),
	}, nil
}

func DecodeGRPCResponseCommentsByAuthorID(_ context.Context, grpcResp interface{}) (interface{}, error) {
	resp := grpcResp.(*pb.ResponseCommentsByAuthorID)
	return Response{
		Data: resp.Comments,
		Err:  util.Str2Err(resp.Err),
	}, nil
}

func DecodeGRPCResponseDeleteCommentByID(_ context.Context, grpcResp interface{}) (interface{}, error) {
	resp := grpcResp.(*pb.ResponseDeleteByID)
	return Response{
		Data: resp.Ok,
		Err:  util.Str2Err(resp.Err),
	}, nil
}
