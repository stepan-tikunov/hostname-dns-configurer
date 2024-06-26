package api

import (
	"net"
	"strconv"

	api "github.com/stepan-tikunov/hostname-dns-configurer/api/gen/go/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	api.HostnameServiceClient
	api.DnsServiceClient
	grpcConn *grpc.ClientConn
}

func (c *Client) Close() {
	c.grpcConn.Close()
}

func NewClient(hostname string, port int) (*Client, error) {
	addr := net.JoinHostPort(hostname, strconv.Itoa(port))
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := &Client{
		HostnameServiceClient: api.NewHostnameServiceClient(conn),
		DnsServiceClient:      api.NewDnsServiceClient(conn),
		grpcConn:              conn,
	}

	return client, nil
}
