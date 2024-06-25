//go:build linux

package hostname

import (
	"context"
	"syscall"

	api "github.com/stepan-tikunov/hostname-dns-configurer/api/gen/go/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"modernc.org/libc/limits"
)

func (s *service) SetHostname(_ context.Context, msg *api.HostnameMessage) (*api.HostnameMessage, error) {
	buf := []byte(msg.GetHostname())
	if len(buf) > limits.HOST_NAME_MAX {
		return nil, status.Errorf(codes.InvalidArgument, "name too long")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	err := syscall.Sethostname(buf)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return msg, err
}
