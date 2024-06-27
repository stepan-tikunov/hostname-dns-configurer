package grpc

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
	hostPort string
}

func (c *Client) Connect() error {
	conn, err := grpc.NewClient(c.hostPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	c.grpcConn = conn

	c.HostnameServiceClient = api.NewHostnameServiceClient(c.grpcConn)
	c.DnsServiceClient = api.NewDnsServiceClient(c.grpcConn)

	return nil
}

func NewClient(hostname string, port int) *Client {
	hp := net.JoinHostPort(hostname, strconv.Itoa(port))

	client := &Client{
		hostPort: hp,
	}

	return client
}
