package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vetcher/comments-msv/service/pb"
)

func DecodeRequestOnlyWithID(_ context.Context, r *http.Request) (interface{}, error) {
	var request RequestOnlyWithID
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

func DecodeGRPCRequestWithID(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.RequestWithID)
	return RequestOnlyWithID{
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
