package api

import (
	"fmt"
	"github.com/stepan-tikunov/hostname-dns-configurer/service/internal/api/dns"
	"github.com/stepan-tikunov/hostname-dns-configurer/service/internal/api/hostname"
	"log/slog"
	"net"

	api "github.com/stepan-tikunov/hostname-dns-configurer/api/gen/go/api/v1"
	"google.golang.org/grpc"
)

func RunService(port int) chan error {
	e := make(chan error, 1)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		e <- err
		return e
	}

	s := grpc.NewServer()
	api.RegisterHostnameServiceServer(s, hostname.NewService())
	api.RegisterDnsServiceServer(s, dns.NewService())

	go func() {
		err = s.Serve(lis)
		e <- err
	}()

	slog.Info("gRPC endpoint launched", slog.Int("port", port))

	return e
}
