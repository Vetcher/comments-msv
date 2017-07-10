package service

import (
	"context"
	"time"

	"log"

	"github.com/go-kit/kit/endpoint"
)

type Middleware func(endpoint.Endpoint) endpoint.Endpoint

func TransportLoggingMiddleware(logger *log.Logger) Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			start := time.Now()
			logger.Printf("%v", request)
			defer logger.Printf("time: %v", time.Since(start))
			return next(ctx, request)
		}
	}
}

func ErrorLoggingMiddleware(logger *log.Logger) Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			d, err := next(ctx, request)
			if err != nil {
				logger.Printf("Internal Server Error: %v", err)
			}
			return d, err
		}
	}
}