package api

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	api "github.com/stepan-tikunov/hostname-dns-configurer/api/gen/go/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func handleStreamError(_ context.Context, err error) *status.Status {
	code := codes.Internal
	msg := "unexpected error"
	if s, ok := status.FromError(err); ok {
		code = s.Code()
		msg = s.Message()
	}
	return status.New(code, msg)
}

func RunGateway(httpPort, grpcPort int) chan error {
	e := make(chan error, 1)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(
		runtime.WithStreamErrorHandler(handleStreamError),
	)
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	endpoint := fmt.Sprintf("127.0.0.1:%d", grpcPort)

	err := api.RegisterHostnameServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		e <- err
		return e
	}

	err = api.RegisterDnsServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		e <- err
		return e
	}

	slog.Info("gRPC gateway REST API launched",
		slog.Int("grpcPort", grpcPort),
		slog.Int("httpPort", httpPort),
	)

	go func() {
		const readTimeout = 3 * time.Second
		server := &http.Server{
			Addr:              fmt.Sprintf("127.0.0.1:%d", httpPort),
			Handler:           mux,
			ReadHeaderTimeout: readTimeout,
		}
		err = server.ListenAndServe()
		e <- err
	}()

	return e
}
