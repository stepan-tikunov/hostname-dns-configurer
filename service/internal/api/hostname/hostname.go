package hostname

import (
	"context"
	"log/slog"
	"os"
	"sync"

	api "github.com/stepan-tikunov/hostname-dns-configurer/api/gen/go/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type service struct {
	api.UnimplementedHostnameServiceServer
	mu *sync.RWMutex
}

func (s *service) GetHostname(_ context.Context, _ *emptypb.Empty) (*api.HostnameMessage, error) {
	slog.Info("Incoming GetHostname() request")

	s.mu.RLock()
	defer s.mu.RUnlock()

	h, err := os.Hostname()
	if err != nil {
		slog.Error("Failed to get hostname", slog.Any("error", err))
		return nil, status.Errorf(codes.Internal, "failed to get hostname: %s", err)
	}

	return &api.HostnameMessage{Hostname: h}, nil
}

func NewService() api.HostnameServiceServer {
	return &service{mu: &sync.RWMutex{}}
}
