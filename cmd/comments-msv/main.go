package main

import (
	"net/http"
	"os"

	"log"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/vetcher/comments-msv/models"
	"github.com/vetcher/comments-msv/service"
)

func main() {
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
	logger.Println("Serve :8081")
	logger.Println(http.ListenAndServe(":8081", nil))
}
