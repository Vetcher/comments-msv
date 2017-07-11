package service

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/vetcher/comments-msv/models"
	"github.com/vetcher/comments-msv/service/pb"
)

func ServeJSON(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func EncodeGRPCResponseComment(_ context.Context, svcResp interface{}) (interface{}, error) {
	resp := svcResp.(*JsonResponse)
	comment := resp.Data.(*models.Comment)
	return &pb.ResponseComment{
		Data: &pb.Comment{
			Id:       uint32(comment.ID),
			Text:     comment.Text,
			AuthorId: uint32(comment.AuthorID),
		},
		Err: resp.Err,
	}, nil
}

func EncodeGRPCResponseBool(_ context.Context, svcResp interface{}) (interface{}, error) {
	resp := svcResp.(*JsonResponse)
	ok := resp.Data.(bool)
	return &pb.ResponseWithBool{
		Ok:  ok,
		Err: resp.Err,
	}, nil
}

// можно как-то иначе сделать?
func ConvDBCommentsToPBComments(coms []*models.Comment) []*pb.Comment {
	var converted []*pb.Comment
	for _, c := range coms {
		converted = append(converted, &pb.Comment{
			Id:       uint32(c.ID),
			AuthorId: uint32(c.AuthorID),
			Text:     c.Text,
		})
	}
	return converted
}

func EncodeGRPCResponseCommentsByAuthorID(_ context.Context, svcResp interface{}) (interface{}, error) {
	resp := svcResp.(*JsonResponse)
	return pb.ResponseCommentsByAuthorID{
		Comments: ConvDBCommentsToPBComments(resp.Data.([]*models.Comment)),
		Err:      resp.Err,
	}, nil
}
