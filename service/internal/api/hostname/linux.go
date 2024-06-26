//go:build linux

package hostname

import (
	"context"
	"log/slog"
	"syscall"

	api "github.com/stepan-tikunov/hostname-dns-configurer/api/gen/go/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"modernc.org/libc/limits"
)

func (s *service) SetHostname(_ context.Context, msg *api.HostnameMessage) (*api.HostnameMessage, error) {
	buf := []byte(msg.GetHostname())
	if len(buf) > limits.HOST_NAME_MAX {
		slog.Error("Failed to set hostname: it is too long", slog.Int("limit", limits.HOST_NAME_MAX))
		return nil, status.Errorf(codes.InvalidArgument, "name too long")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	err := syscall.Sethostname(buf)
	if err != nil {
		slog.Error("Failed to set hostname", slog.Any("error", err))
		return nil, status.Errorf(codes.Internal, "failed to set hostname: %s", err)
	}

	slog.Info("Updated hostname", slog.String("hostname", msg.GetHostname()))

	return msg, nil
}
