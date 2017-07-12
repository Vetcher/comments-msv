package service

import (
	"testing"

	"github.com/vetcher/comments-msv/client"
	"google.golang.org/grpc"
)

const GRPCComment string = "I'm testing GRPC"
const Target string = "localhost:10000"

func TestEndpoints_PostComment(t *testing.T) {
	conn, err := grpc.Dial(Target, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	grpcClient := client.NewClient(conn)
	comment, err := grpcClient.PostComment(1, GRPCComment)
	if err != nil {
		t.Fatalf("client error: %v", err)
	}
	if comment == nil || !(comment.Text == GRPCComment && comment.AuthorID == 1 && comment.ID == 1) {
		t.Fatalf("wrong comment: %v", comment)
	}
}

func TestEndpoints_GetCommentByID(t *testing.T) {
	conn, err := grpc.Dial(Target, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	grpcClient := client.NewClient(conn)
	comment, err := grpcClient.GetCommentByID(1)
	if err != nil {
		t.Fatalf("client error: %v", err)
	}
	if comment == nil || !(comment.ID == 1 && comment.Text == GRPCComment && comment.AuthorID == 1) {
		t.Fatalf("comment not ideal: %v", comment)
	}
}

func TestEndpoints_GetCommentsByAuthorID(t *testing.T) {
	conn, err := grpc.Dial(Target, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	grpcClient := client.NewClient(conn)
	comments, err := grpcClient.GetCommentsByAuthorID(1)
	if err != nil {
		t.Fatalf("client error: %v", err)
	}
	if comments == nil || !(len(comments) >= 1 && comments[0].AuthorID == 1) {
		t.Fatalf("comments not ideal: %v", comments)
	}
}

func TestEndpoints_DeleteCommentByID(t *testing.T) {
	conn, err := grpc.Dial(Target, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	grpcClient := client.NewClient(conn)
	ok, err := grpcClient.DeleteCommentByID(1)
	if err != nil {
		t.Fatalf("client error: %v", err)
	}
	if !ok {
		t.Fatalf("Response schould be `true`")
	}
	_, err = grpcClient.GetCommentByID(1)
	if err == nil || err.Error() != "database error: record not found" {
		t.Fatalf("Have: `%v`\nExpected: `%v`", err, "database error: record not found")
	}
}
