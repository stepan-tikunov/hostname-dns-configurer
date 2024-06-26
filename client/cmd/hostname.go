package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	apiv1 "github.com/stepan-tikunov/hostname-dns-configurer/api/gen/go/api/v1"
	"github.com/stepan-tikunov/hostname-dns-configurer/client/internal/api"
)

var hostnameCmd = &cobra.Command{
	Use:   "hostname",
	Short: "Configure hostname",
	Long:  "Get or set service's hostname",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		return cmd.Help()
	},
}

var getHostnameCmd = &cobra.Command{
	Use:   "get",
	Short: "Get hostname",
	Long:  "Get service's hostname",
	RunE: func(_ *cobra.Command, _ []string) error {
		cl, err := api.NewClient(serviceHost, servicePort)
		if err != nil {
			return err
		}

		resp, err := cl.GetHostname(context.Background(), nil)
		if err != nil {
			return err
		}

		fmt.Println(resp.GetHostname())

		return nil
	},
}

var setHostnameCmd = &cobra.Command{
	Use:   "set [hostname]",
	Short: "Set hostname",
	Long:  "Set service's hostname",
	Args:  cobra.ExactArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		cl, err := api.NewClient(serviceHost, servicePort)
		if err != nil {
			return err
		}

		req := apiv1.HostnameMessage{Hostname: args[0]}
		resp, err := cl.SetHostname(context.Background(), &req)
		if err != nil {
			return err
		}

		// Separate output to allow for piping useful info to other commands.
		fmt.Fprintf(os.Stderr, "Hostname updated: ")
		fmt.Println(resp.GetHostname())

		return nil
	},
}

func init() {
	rootCmd.AddCommand(hostnameCmd)
	hostnameCmd.AddCommand(getHostnameCmd)
	hostnameCmd.AddCommand(setHostnameCmd)
}
