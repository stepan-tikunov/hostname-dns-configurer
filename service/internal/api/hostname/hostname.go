package hostname

import (
	"context"
	"os"
	"sync"

	api "github.com/stepan-tikunov/hostname-dns-configurer/api/gen/go/api/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type service struct {
	api.UnimplementedHostnameServiceServer
	mu *sync.RWMutex
}

func (s *service) GetHostname(_ context.Context, _ *emptypb.Empty) (*api.HostnameMessage, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	h, err := os.Hostname()

	return &api.HostnameMessage{Hostname: h}, err
}

func NewService() api.HostnameServiceServer {
	return &service{mu: &sync.RWMutex{}}
}
