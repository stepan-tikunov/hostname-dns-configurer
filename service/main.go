package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/stepan-tikunov/hostname-dns-configurer/service/internal/api"
)

func exitError(msg string, err error) {
	slog.Error(msg, slog.Any("error", err))
	os.Exit(1)
}

func main() {
	grpcPort := flag.Int("api", 9000, "gRPC endpoint port")
	httpPort := flag.Int("gateway", 8080, "gRPC-gateway port")

	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	serviceErr := api.RunService(*grpcPort)
	gatewayErr := api.RunGateway(*httpPort, *grpcPort)

	select {
	case err := <-serviceErr:
		exitError("gRPC service stopped", err)
	case err := <-gatewayErr:
		exitError("gRPC gateway stopped", err)
	}
}
