package main

import (
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/vetcher/comments-msv/service"
)

func main() {
	svc := service.CommentsService{}
	svc.Init()
	handlerGetByID := httptransport.NewServer(
		service.GetCommentEndpoint(svc),
		service.DecodeRequestOnlyWithId,
		service.ServeJSON,
	)
	handlerPost := httptransport.NewServer(
		service.PostCommentEndpoint(svc),
		service.DecodeRequestOnlyWithId,
		service.ServeJSON,
	)
	handlerDelete := httptransport.NewServer(
		service.DeleteCommentEndpoint(svc),
		service.DecodeRequestPostComment,
		service.ServeJSON,
	)
	handlerGetForAuthor := httptransport.NewServer(
		service.GetCommentsForUserEndpoint(svc),
		service.DecodeRequestOnlyWithId,
		service.ServeJSON,
	)
	http.Handle("/comment/get", handlerGetByID)
	http.Handle("/comment/post", handlerPost)
	http.Handle("/comment/del", handlerDelete)
	http.Handle("/comments/author", handlerGetForAuthor)
	log.Println("Serve :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
