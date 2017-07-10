package main

import (
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/vetcher/comments-msv/service"
)

func main() {
	svc := service.NewCommentService()
	handlerGetByID := httptransport.NewServer(
		service.GetCommentEndpoint(svc),
		service.DecodeRequestOnlyWithID,
		service.ServeJSON,
	)
	handlerPost := httptransport.NewServer(
		service.PostCommentEndpoint(svc),
		service.DecodeRequestPostComment,
		service.ServeJSON,
	)
	handlerDeleteByID := httptransport.NewServer(
		service.DeleteCommentEndpoint(svc),
		service.DecodeRequestOnlyWithID,
		service.ServeJSON,
	)
	handlerGetByAuthorID := httptransport.NewServer(
		service.GetCommentsByAuthorIDEndpoint(svc),
		service.DecodeRequestOnlyWithID,
		service.ServeJSON,
	)
	http.Handle("/comment/get", handlerGetByID)
	http.Handle("/comment/post", handlerPost)
	http.Handle("/comment/del", handlerDeleteByID)
	http.Handle("/comments/author", handlerGetByAuthorID)
	log.Println("Serve :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
