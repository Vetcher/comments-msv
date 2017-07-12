package main

import (
	"log"
	"net/http"
	"os"

	"net"

	"flag"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/vetcher/comments-msv/models"
	"github.com/vetcher/comments-msv/service"
	"github.com/vetcher/comments-msv/service/pb"
	"github.com/vetcher/comments-msv/transport"
	"google.golang.org/grpc"
)

func startHTTPService(db *models.Database) {
	logger := log.New(os.Stdout, "http | ", log.Lshortfile|log.Ltime)
	svc := service.NewCommentService(db)
	handlerGetByID := httptransport.NewServer(
		service.TransportLoggingMiddleware(logger)(service.GetCommentEndpoint(svc)),
		service.DecodeRequestOnlyWithID,
		service.ServeJSON,
	)
	handlerPost := httptransport.NewServer(
		service.TransportLoggingMiddleware(logger)(service.PostCommentEndpoint(svc)),
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

func startGRPCService(db *models.Database) {
	logger := log.New(os.Stdout, "grpc | ", log.Lshortfile|log.Ltime)
	lis, err := net.Listen("tcp", ":10000")
	if err != nil {
		logger.Fatalf("Can't listen: %v", err)
	}
	svc := service.NewCommentService(db)
	handlerGetByID := grpctransport.NewServer(
		service.TransportLoggingMiddleware(logger)(service.GetCommentEndpoint(svc)),
		service.DecodeGRPCRequestWithID,
		service.EncodeGRPCResponseComment,
	)
	handlerPost := grpctransport.NewServer(
		service.TransportLoggingMiddleware(logger)(service.PostCommentEndpoint(svc)),
		service.DecodeGRPCRequestPostComment,
		service.EncodeGRPCResponseComment,
	)
	handlerDeleteByID := grpctransport.NewServer(
		service.TransportLoggingMiddleware(logger)(service.DeleteCommentEndpoint(svc)),
		service.DecodeGRPCRequestWithID,
		service.EncodeGRPCResponseBool,
	)
	handlerGetByAuthorID := grpctransport.NewServer(
		service.TransportLoggingMiddleware(logger)(service.GetCommentsByAuthorIDEndpoint(svc)),
		service.DecodeGRPCRequestWithID,
		service.EncodeGRPCResponseCommentsByAuthorID,
	)
	server := transport.GRPCServer{
		GetCommentByIDHandler:        handlerGetByID,
		PostCommentHandler:           handlerPost,
		DeleteCommentByIDHandler:     handlerDeleteByID,
		GetCommentsByAuthorIDHandler: handlerGetByAuthorID,
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCommentSVCServer(grpcServer, server)
	logger.Println("Serve :10000")
	logger.Println(grpcServer.Serve(lis))
}

func main() {
	db := models.NewDatabase()
	flag.Parse()
	go startHTTPService(db)
	startGRPCService(db)
}
