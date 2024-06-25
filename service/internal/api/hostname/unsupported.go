//go:build !linux

package hostname

import (
	"context"
	"log/slog"

	api "github.com/stepan-tikunov/hostname-dns-configurer/api/gen/go/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) SetHostname(_ context.Context, msg *api.HostnameMessage) (*api.HostnameMessage, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	err := status.Error(codes.Unimplemented, "setting the hostname is not supported")

	slog.Error("Error processing incoming request",
		slog.String("rpc", "HostnameService.SetHostname()"),
		slog.Any("request", msg),
		slog.Any("error", err),
	)

	return nil, err
}
