package interceptors

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func TimingInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	start := time.Now()

	defer func() {
		duration := time.Since(start)

		log.Printf("handler %s took %v ms", info.FullMethod, duration)
	}()

	return handler(ctx, req)
}