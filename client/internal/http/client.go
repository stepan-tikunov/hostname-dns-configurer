package http

import (
	"context"
	"net"
	"strconv"

	api "github.com/stepan-tikunov/hostname-dns-configurer/api/gen/go/api/v1"
	"github.com/stepan-tikunov/hostname-dns-configurer/api/gen/go/client"
	"github.com/stepan-tikunov/hostname-dns-configurer/api/gen/go/client/dns_service"
	"github.com/stepan-tikunov/hostname-dns-configurer/api/gen/go/client/hostname_service"
	"github.com/stepan-tikunov/hostname-dns-configurer/api/gen/go/swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Client struct {
	http *client.APIV1DNSProto
}

type errorWithRPCStatusPayload interface {
	GetPayload() *swagger.RPCStatus
}

func extractStatus(err error) error {
	if e, ok := err.(errorWithRPCStatusPayload); ok {
		p := e.GetPayload()
		err = status.Error(codes.Code(p.Code), p.Message)
	}

	return err
}

func (c *Client) SetHostname(_ context.Context, in *api.HostnameMessage, _ ...grpc.CallOption) (
	*api.HostnameMessage, error,
) {
	params := hostname_service.NewHostnameServiceSetHostnameParams()
	params.SetBody(&swagger.V1HostnameMessage{
		Hostname: in.GetHostname(),
	})

	resp, err := c.http.HostnameService.HostnameServiceSetHostname(params)
	if err != nil {
		return nil, extractStatus(err)
	}

	msg := &api.HostnameMessage{
		Hostname: resp.GetPayload().Hostname,
	}

	return msg, nil
}

func (c *Client) GetHostname(_ context.Context, _ *emptypb.Empty, _ ...grpc.CallOption) (
	*api.HostnameMessage, error,
) {
	params := hostname_service.NewHostnameServiceGetHostnameParams()

	resp, err := c.http.HostnameService.HostnameServiceGetHostname(params)
	if err != nil {
		return nil, extractStatus(err)
	}

	msg := &api.HostnameMessage{
		Hostname: resp.GetPayload().Hostname,
	}

	return msg, nil
}

func (c *Client) GetNameserverList(_ context.Context, _ *emptypb.Empty, _ ...grpc.CallOption) (
	*api.NameserverList, error,
) {
	params := dns_service.NewDNSServiceGetNameserverListParams()

	resp, err := c.http.DNSService.DNSServiceGetNameserverList(params)
	if err != nil {
		return nil, extractStatus(err)
	}
	payload := resp.GetPayload()
	servers := make([]*api.Nameserver, 0, len(payload.Servers))
	for _, server := range payload.Servers {
		servers = append(servers, &api.Nameserver{
			Index:   server.Index,
			Address: server.Address,
		})
	}

	msg := &api.NameserverList{
		Servers:  servers,
		Checksum: uint32(payload.Checksum),
	}

	return msg, nil
}

func (c *Client) GetNameserverAt(_ context.Context, _ *api.GetNameserverRequest, _ ...grpc.CallOption) (
	*api.NameserverResponse, error,
) {
	params := dns_service.NewDNSServiceGetNameserverAtParams()

	resp, err := c.http.DNSService.DNSServiceGetNameserverAt(params)
	if err != nil {
		return nil, extractStatus(err)
	}

	payload := resp.GetPayload()
	msg := &api.NameserverResponse{
		Server: &api.Nameserver{
			Index:   payload.Server.Index,
			Address: payload.Server.Address,
		},
		Checksum: uint32(payload.Checksum),
	}

	return msg, nil
}

func (c *Client) CreateNameserver(_ context.Context, in *api.CreateNameserverRequest, _ ...grpc.CallOption) (
	*api.NameserverResponse, error,
) {
	params := dns_service.NewDNSServiceCreateNameserverParams()

	params.Index = in.Index
	params.Address = in.GetAddress()
	if in.Checksum != nil {
		checksum := int64(in.GetChecksum())
		params.Checksum = &checksum
	}

	resp, err := c.http.DNSService.DNSServiceCreateNameserver(params)
	if err != nil {
		return nil, extractStatus(err)
	}

	payload := resp.GetPayload()
	msg := &api.NameserverResponse{
		Server: &api.Nameserver{
			Index:   payload.Server.Index,
			Address: payload.Server.Address,
		},
		Checksum: uint32(payload.Checksum),
	}

	return msg, nil
}

func (c *Client) DeleteNameserver(_ context.Context, in *api.DeleteNameserverRequest, _ ...grpc.CallOption) (
	*api.NameserverResponse, error,
) {
	checksum := int64(in.GetChecksum())

	params := dns_service.NewDNSServiceDeleteNameserverParams()
	params.Checksum = &checksum
	params.Index = in.GetIndex()

	resp, err := c.http.DNSService.DNSServiceDeleteNameserver(params)
	if err != nil {
		return nil, extractStatus(err)
	}

	payload := resp.GetPayload()
	msg := &api.NameserverResponse{
		Server: &api.Nameserver{
			Index:   payload.Server.Index,
			Address: payload.Server.Address,
		},
		Checksum: uint32(payload.Checksum),
	}

	return msg, nil
}

func (c *Client) Connect() error {
	return nil
}

func NewClient(hostname string, port int) *Client {
	cfg := client.TransportConfig{
		Host:     net.JoinHostPort(hostname, strconv.Itoa(port)),
		BasePath: "/",
		Schemes:  []string{"http"},
	}

	return &Client{
		http: client.NewHTTPClientWithConfig(nil, &cfg),
	}
}
