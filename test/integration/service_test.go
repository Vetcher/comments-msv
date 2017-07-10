package service_test

import (
	"net/http"
	"testing"

	"os"

	"log"

	"bytes"
	"encoding/json"

	"io/ioutil"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/vetcher/comments-msv/models"
	"github.com/vetcher/comments-msv/service"
)

type Comment struct {
	ID       uint   `json:"id,omitempty"`
	Text     string `json:"text"`
	AuthorID uint   `json:"author"`
}

func TestPostComment(t *testing.T) {
	const CommentText = "Testing Comment"
	comment := Comment{
		Text:     CommentText,
		AuthorID: 1,
	}
	b, err := json.Marshal(&comment)
	if err != nil {
		t.Fatalf("Marshal to JSON error: %v", err)
	}
	buf := bytes.NewBuffer(b)
	request := http.Request{}
	request.Write(buf)
	resp, err := http.NewRequest("GET", "http://localhost:8081/comment/post", buf)
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("reading response Body error: %v", err)
	}
	var c Comment
	json.Unmarshal(respBody, &c)
	if !(c.AuthorID == 1 && c.Text == CommentText) {
		t.Fatalf("%v != %v", c, Comment{Text: CommentText, AuthorID: 1})
	}
	log.Println(c)
}

func TestMain(m *testing.M) {
	logger := log.New(os.Stdout, "service: ", log.Lshortfile|log.Ltime)
	db := models.NewDatabase()
	svc := service.NewCommentService(db)
	handlerGetByID := httptransport.NewServer(
		service.TransportLoggingMiddleware(logger)(service.GetCommentEndpoint(svc)),
		service.DecodeRequestOnlyWithID,
		service.ServeJSON,
	)
	postCommentEndpoint := service.PostCommentEndpoint(svc)
	postCommentEndpoint = service.ErrorLoggingMiddleware(logger)(postCommentEndpoint)
	postCommentEndpoint = service.TransportLoggingMiddleware(logger)(postCommentEndpoint)
	handlerPost := httptransport.NewServer(
		postCommentEndpoint,
		service.DecodeRequestPostComment,
		service.ServeJSON,
	)
	handlerDeleteByID := httptransport.NewServer(
		service.TransportLoggingMiddleware(logger)(service.DeleteCommentEndpoint(svc)),
		service.DecodeRequestOnlyWithID,
		service.ServeJSON,
	)
	handlerGetByAuthorID := httptransport.NewServer(
		service.TransportLoggingMiddleware(logger)(service.GetCommentsByAuthorIDEndpoint(svc)),
		service.DecodeRequestOnlyWithID,
		service.ServeJSON,
	)
	http.Handle("/comment/get", handlerGetByID)
	http.Handle("/comment/post", handlerPost)
	http.Handle("/comment/del", handlerDeleteByID)
	http.Handle("/comments/author", handlerGetByAuthorID)
	go http.ListenAndServe(":8081", nil)
	os.Exit(m.Run())
}
