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

const POSTGRES string = "postgres"

func parseDBConfig() *models.DBConfig {
	conf := models.DBConfig{
		User:     flag.String("user", POSTGRES, "User"),
		Password: flag.String("password", POSTGRES, "User's password"),
		DBName:   flag.String("db", POSTGRES, "Name of database"),
		Port:     flag.Uint("port", 5432, "Postgres port"),
		Host:     flag.String("host", "localhost", "Address of server"),
	}
	flag.Parse()
	return &conf
}

func startHTTPService(db *models.Database) error {
	logger := log.New(os.Stdout, "http | ", log.Ltime)
	svc := service.NewCommentService(db)
	handlerGetByID := httptransport.NewServer(
		service.TransportLoggingMiddleware(logger)(service.GetCommentEndpoint(svc)),
		service.DecodeRequestGetCommentByID,
		service.ServeJSON,
	)
	handlerPost := httptransport.NewServer(
		service.TransportLoggingMiddleware(logger)(service.PostCommentEndpoint(svc)),
		service.DecodeRequestPostComment,
		service.ServeJSON,
	)
	handlerDeleteByID := httptransport.NewServer(
		service.TransportLoggingMiddleware(logger)(service.DeleteCommentEndpoint(svc)),
		service.DecodeRequestDeleteCommentByID,
		service.ServeJSON,
	)
	handlerGetByAuthorID := httptransport.NewServer(
		service.TransportLoggingMiddleware(logger)(service.GetCommentsByAuthorIDEndpoint(svc)),
		service.DecodeRequestGetCommentByAuthorID,
		service.ServeJSON,
	)
	http.Handle("/comment/get", handlerGetByID)
	http.Handle("/comment/post", handlerPost)
	http.Handle("/comment/del", handlerDeleteByID)
	http.Handle("/comments/author", handlerGetByAuthorID)
	logger.Println("Serve :8081")
	return http.ListenAndServe(":8081", nil)
}

func startGRPCService(db *models.Database) error {
	logger := log.New(os.Stdout, "grpc | ", log.Ltime)
	lis, err := net.Listen("tcp", ":10000")
	if err != nil {
		logger.Fatalf("Can't listen: %v", err)
	}
	svc := service.NewCommentService(db)
	handlerGetByID := grpctransport.NewServer(
		service.TransportLoggingMiddleware(logger)(service.GetCommentEndpoint(svc)),
		service.DecodeGRPCRequestGetCommentByID,
		service.EncodeGRPCResponseGetCommentByID,
	)
	handlerPost := grpctransport.NewServer(
		service.TransportLoggingMiddleware(logger)(service.PostCommentEndpoint(svc)),
		service.DecodeGRPCRequestPostComment,
		service.EncodeGRPCResponsePostComment,
	)
	handlerDeleteByID := grpctransport.NewServer(
		service.TransportLoggingMiddleware(logger)(service.DeleteCommentEndpoint(svc)),
		service.DecodeGRPCRequestDeleteCommentByID,
		service.EncodeGRPCResponseDeleteCommentByID,
	)
	handlerGetByAuthorID := grpctransport.NewServer(
		service.TransportLoggingMiddleware(logger)(service.GetCommentsByAuthorIDEndpoint(svc)),
		service.DecodeGRPCRequestGetCommentByAuthorID,
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
	return grpcServer.Serve(lis)
}

func main() {
	db := models.NewDatabase(parseDBConfig())
	errChan := make(chan error)
	go func() {
		errChan <- startHTTPService(db)
	}()
	go func() {
		errChan <- startGRPCService(db)
	}()
	log.Fatal(<-errChan)
}
