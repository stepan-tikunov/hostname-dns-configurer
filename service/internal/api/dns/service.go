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
	rc := dns.GetResolvConfInstance()
	ns, checksum, err := rc.GetNameservers()

	if err != nil {
		slog.Error("failed to get nameservers", slog.Any("error", err))
		return nil, status.Errorf(codes.Internal, "failed to get nameservers: %v", err)
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

	return list, err
}

func (s service) GetNameserverAt(
	_ context.Context,
	request *api.GetNameserverRequest,
) (*api.NameserverResponse, error) {
	rc := dns.GetResolvConfInstance()
	index := request.GetIndex()
	n, checksum, err := rc.GetNameserverAt(int(index))
	if err != nil {
		slog.Error("failed to get nameserver", slog.Int("index", int(index)), slog.Any("error", err))
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
		return nil, status.Errorf(codes.Internal, "failed to create nameserver: %v", err)
	}

	response := &api.NameserverResponse{
		Server: &api.Nameserver{
			Index:   int32(index),
			Address: nameserver,
		},
		Checksum: checksum,
	}

	return response, nil
}

func (s service) CreateNameserver(
	_ context.Context,
	request *api.CreateNameserverRequest,
) (*api.NameserverResponse, error) {
	if request.Index == nil {
		return s.createNameserverLast(request)
	}

	rc := dns.GetResolvConfInstance()

	checksum := request.GetChecksum()
	index := int(request.GetIndex())
	nameserver := request.GetAddress()

	checksum, err := rc.CreateNameserverAt(checksum, index, nameserver)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create nameserver: %v", err)
	}

	response := &api.NameserverResponse{
		Server: &api.Nameserver{
			Index:   int32(index),
			Address: nameserver,
		},
		Checksum: checksum,
	}

	return response, nil
}

func (s service) DeleteNameserver(
	_ context.Context,
	request *api.DeleteNameserverRequest,
) (*api.NameserverResponse, error) {
	rc := dns.GetResolvConfInstance()

	checksum := request.GetChecksum()
	index := int(request.GetIndex())

	nameserver, checksum, err := rc.DeleteNameserverAt(checksum, index)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete nameserver: %v", err)
	}

	response := &api.NameserverResponse{
		Server: &api.Nameserver{
			Index:   int32(index),
			Address: nameserver,
		},
		Checksum: checksum,
	}

	return response, nil
}

func NewService() api.DnsServiceServer {
	return &service{}
}
