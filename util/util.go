package util

import (
	"errors"

	"github.com/vetcher/comments-msv/models"
	"github.com/vetcher/comments-msv/service/pb"
)

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

func Str2Err(str string) error {
	if str == "" {
		return nil
	}
	return errors.New(str)
}
