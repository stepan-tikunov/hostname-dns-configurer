package dns

import (
	"context"
	"log/slog"

	api "github.com/stepan-tikunov/hostname-dns-configurer/api/gen/go/api/v1"
	"github.com/stepan-tikunov/hostname-dns-configurer/service/internal/dns"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type service struct {
	api.UnimplementedDnsServiceServer
}

func (s service) GetNameserverList(_ context.Context, _ *emptypb.Empty) (*api.NameserverList, error) {
	slog.Info("Incoming GetNameserverList() request")
	rc := dns.GetResolvConfInstance()
	ns, checksum, err := rc.GetNameservers()

	if err != nil {
		slog.Error("Failed to get nameservers", slog.Any("error", err))
		return nil, status.Errorf(codes.Internal, "failed to get nameservers: %s", err)
	}

	servers := make([]*api.Nameserver, 0, len(ns))
	for i, n := range ns {
		servers = append(servers, &api.Nameserver{
			Index:   int32(i),
			Address: n,
		})
	}

	list := &api.NameserverList{
		Servers:  servers,
		Checksum: checksum,
	}

	return list, nil
}

func (s service) GetNameserverAt(
	_ context.Context,
	request *api.GetNameserverRequest,
) (*api.NameserverResponse, error) {
	slog.Info("Incoming GetNameserverAt() request", slog.Any("request", request))
	rc := dns.GetResolvConfInstance()
	index := request.GetIndex()
	n, checksum, err := rc.GetNameserverAt(int(index))
	if err != nil {
		slog.Error("Failed to get nameserver", slog.Int("index", int(index)), slog.Any("error", err))
		return nil, err
	}

	resp := &api.NameserverResponse{
		Server: &api.Nameserver{
			Index:   index,
			Address: n,
		},
		Checksum: checksum,
	}

	return resp, nil
}

func (s service) createNameserverLast(request *api.CreateNameserverRequest) (*api.NameserverResponse, error) {
	rc := dns.GetResolvConfInstance()

	nameserver := request.GetAddress()

	index, checksum, err := rc.CreateNameserverLast(nameserver)
	if err != nil {
		slog.Error("Failed to create nameserver", slog.Any("error", err))
		return nil, status.Errorf(codes.Internal, "failed to create nameserver: %s", err)
	}

	response := &api.NameserverResponse{
		Server: &api.Nameserver{
			Index:   int32(index),
			Address: nameserver,
		},
		Checksum: checksum,
	}

	slog.Info("Created nameserver", slog.Int("index", index), slog.String("address", nameserver))

	return response, nil
}

func (s service) CreateNameserver(
	_ context.Context,
	request *api.CreateNameserverRequest,
) (*api.NameserverResponse, error) {
	slog.Info("Incoming CreateNameserver() request", slog.Any("request", request))

	if request.Index == nil {
		return s.createNameserverLast(request)
	}

	rc := dns.GetResolvConfInstance()

	checksum := request.GetChecksum()
	index := int(request.GetIndex())
	nameserver := request.GetAddress()

	checksum, err := rc.CreateNameserverAt(checksum, index, nameserver)
	if err != nil {
		slog.Error("Failed to create nameserver", slog.Any("error", err))
		return nil, status.Errorf(codes.Internal, "failed to create nameserver: %s", err)
	}

	response := &api.NameserverResponse{
		Server: &api.Nameserver{
			Index:   int32(index),
			Address: nameserver,
		},
		Checksum: checksum,
	}

	slog.Info("Created nameserver", slog.Int("index", index), slog.String("address", nameserver))

	return response, nil
}

func (s service) DeleteNameserver(
	_ context.Context,
	request *api.DeleteNameserverRequest,
) (*api.NameserverResponse, error) {
	slog.Info("Incoming DeleteNameserver() request", slog.Any("request", request))

	rc := dns.GetResolvConfInstance()

	checksum := request.GetChecksum()
	index := int(request.GetIndex())

	nameserver, checksum, err := rc.DeleteNameserverAt(checksum, index)
	if err != nil {
		slog.Error("Failed to delete nameserver", slog.Any("error", err))
		return nil, status.Errorf(codes.Internal, "failed to delete nameserver: %s", err)
	}

	response := &api.NameserverResponse{
		Server: &api.Nameserver{
			Index:   int32(index),
			Address: nameserver,
		},
		Checksum: checksum,
	}

	slog.Info("Deleted nameserver", slog.Int("index", index), slog.String("address", nameserver))

	return response, nil
}

func NewService() api.DnsServiceServer {
	return &service{}
}
