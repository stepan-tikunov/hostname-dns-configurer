package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	api "github.com/stepan-tikunov/hostname-dns-configurer/api/gen/go/api/v1"
	"github.com/stepan-tikunov/hostname-dns-configurer/client/internal/grpc"
	"github.com/stepan-tikunov/hostname-dns-configurer/client/internal/http"
)

func exitError(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

var rootCmd = &cobra.Command{
	Use:           "client",
	Short:         "CLI for hostname & DNS configuration service",
	SilenceUsage:  true,
	SilenceErrors: true,
	Long: "This is the command-line interface for the hostname and DNS\n" +
		"configuration service built with Cobra. You can learn more at\n" +
		"https://github.com/stepan-tikunov/hostname-dns-configurer",
	RunE: func(cmd *cobra.Command, _ []string) error {
		return cmd.Help()
	},
}

type APIClient interface {
	api.HostnameServiceClient
	api.DnsServiceClient

	Connect() error
}

var serviceHost string
var servicePort int
var useHTTP bool

var client APIClient

func init() {
	rootCmd.PersistentFlags().StringVarP(&serviceHost, "host", "H", "127.0.0.1", "service's hostname or IP address")

	const defaultPort = 9000 // To suppress linter
	rootCmd.PersistentFlags().IntVarP(&servicePort, "port", "P", defaultPort, "service's port")

	rootCmd.PersistentFlags().BoolVar(&useHTTP, "http", false, "use HTTP instead of gRPC to connect to service")
}

func Execute() {
	_ = rootCmd.ParseFlags(os.Args[1:])
	if useHTTP {
		client = http.NewClient(serviceHost, servicePort)
	} else {
		client = grpc.NewClient(serviceHost, servicePort)
	}

	if cmd, err := rootCmd.ExecuteC(); err != nil {
		if msg, ok := grpc.ErrorMessage(err); ok {
			exitError(msg)
		}

		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		if cmd != nil {
			_ = cmd.Usage()
		} else {
			_ = rootCmd.Usage()
		}
	}
}
