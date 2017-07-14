package service

import (
	"testing"

	"github.com/vetcher/comments-msv/service"
	client "github.com/vetcher/comments-msv/transport"
	"google.golang.org/grpc"
)

const (
	GRPCComment string = "I'm testing GRPC"
	Target      string = "localhost:10000"
)

var (
	testCommentID uint = 5
	grpcClient    service.CommentService
)

func TestEndpoints_PostComment(t *testing.T) {
	comment, err := grpcClient.PostComment(1, GRPCComment)
	if err != nil {
		t.Fatalf("client error: %v", err)
	}
	if comment == nil || !(comment.Text == GRPCComment && comment.AuthorID == 1) {
		t.Fatalf("wrong comment: %v", comment)
	}
	testCommentID = comment.ID
}

func TestEndpoints_GetCommentByID(t *testing.T) {
	comment, err := grpcClient.GetCommentByID(testCommentID)
	if err != nil {
		t.Fatalf("client error: %v", err)
	}
	if comment == nil || !(comment.ID == testCommentID && comment.Text == GRPCComment && comment.AuthorID == 1) {
		t.Fatalf("comment not ideal: %v", comment)
	}
}

func TestEndpoints_GetCommentsByAuthorID(t *testing.T) {
	comments, err := grpcClient.GetCommentsByAuthorID(1)
	if err != nil {
		t.Fatalf("client error: %v", err)
	}
	if comments == nil || !(len(comments) >= 1 && comments[0].AuthorID == 1) {
		t.Fatalf("comments not ideal: %v", comments)
	}
}

func TestEndpoints_DeleteCommentByID(t *testing.T) {
	ok, err := grpcClient.DeleteCommentByID(testCommentID)
	if err != nil {
		t.Fatalf("client error: %v", err)
	}
	if !ok {
		t.Fatalf("Response schould be `true`")
	}
	_, err = grpcClient.GetCommentByID(testCommentID)
	if err == nil || err.Error() != "database error: record not found" {
		t.Fatalf("Have: `%v`\nExpected: `%v`", err, "database error: record not found")
	}
}

func TestMain(m *testing.M) {
	conn, err := grpc.Dial(Target, grpc.WithInsecure())
	if err == nil {
		grpcClient = client.NewClient(conn)
		defer conn.Close()
		m.Run()
	}
}
