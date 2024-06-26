package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	apiv1 "github.com/stepan-tikunov/hostname-dns-configurer/api/gen/go/api/v1"
	"github.com/stepan-tikunov/hostname-dns-configurer/client/internal/api"
	"github.com/stepan-tikunov/hostname-dns-configurer/client/internal/prompt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "Configure DNS servers list",
	Long:  "Get, add or delete DNS nameservers in service's /etc/resolv.conf file",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		return cmd.Help()
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all DNS servers",
	Long:  "List all DNS servers in service's /etc/resolv.conf file",
	RunE: func(_ *cobra.Command, _ []string) error {
		cl, err := api.NewClient(serviceHost, servicePort)
		if err != nil {
			return err
		}

		resp, err := cl.GetNameserverList(context.Background(), nil)
		if err != nil {
			return err
		}

		fmt.Fprintf(os.Stderr, "Index\tAddress\n")
		for _, ns := range resp.GetServers() {
			fmt.Printf("%d\t%s\n", ns.GetIndex(), ns.GetAddress())
		}

		return nil
	},
}

func getNameserverAndPrompt(msg string, cl *api.Client, index int32) (checksum uint32, err error) {
	resp, err := cl.GetNameserverAt(context.Background(), &apiv1.GetNameserverRequest{Index: index})
	if err != nil {
		return
	}

	server := resp.GetServer()
	if server == nil {
		err = status.Error(codes.DataLoss, "response corrupted")
		return
	}

	fmt.Fprintf(os.Stderr, "Nameserver at index %d: %s\n", index, server.GetAddress())
	if ok := prompt.YesNo(msg); !ok {
		exitError("Aborted")
	}

	checksum = resp.GetChecksum()

	return
}

var addCmd = &cobra.Command{
	Use:   "add [address]",
	Short: "Add a DNS server",
	Long:  "Add a DNS server to service's /etc/resolv.conf file",
	Args:  cobra.ExactArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		cl, err := api.NewClient(serviceHost, servicePort)
		if err != nil {
			return err
		}

		var index *int32
		var checksum *uint32
		if indexFlag >= 0 {
			i := int32(indexFlag)

			var c uint32
			c, err = getNameserverAndPrompt(
				"Do you want to insert new nameserver before it?",
				cl, i,
			)
			if err != nil {
				return err
			}

			checksum = &c
			index = &i
		}

		req := apiv1.CreateNameserverRequest{
			Index:    index,
			Address:  args[0],
			Checksum: checksum,
		}

		resp, err := cl.CreateNameserver(context.Background(), &req)
		if err != nil {
			return err
		}

		ns := resp.GetServer()
		if ns == nil {
			return status.Error(codes.DataLoss, "response corrupted")
		}

		fmt.Fprintf(os.Stderr, "Nameserver added successfully:\n\nIndex\tAddress\n")
		fmt.Printf("%d\t%s\n", ns.GetIndex(), ns.GetAddress())

		return nil
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a DNS server",
	Long:  "Delete a DNS server from service's /etc/resolv.conf file",
	RunE: func(_ *cobra.Command, _ []string) error {
		cl, err := api.NewClient(serviceHost, servicePort)
		if err != nil {
			return err
		}

		i := int32(indexFlag)
		if i < 0 {
			return errors.New("index flag must be set to non-negative value")
		}

		checksum, err := getNameserverAndPrompt(
			"Do you want to delete it?",
			cl, i,
		)
		if err != nil {
			return err
		}

		req := apiv1.DeleteNameserverRequest{
			Index:    i,
			Checksum: checksum,
		}

		resp, err := cl.DeleteNameserver(context.Background(), &req)
		if err != nil {
			return err
		}

		ns := resp.GetServer()
		if ns == nil {
			return status.Error(codes.DataLoss, "response corrupted")
		}

		fmt.Fprintf(os.Stderr, "Nameserver deleted successfully:\n\nIndex\tAddress\n")
		fmt.Printf("%d\t%s\n", ns.GetIndex(), ns.GetAddress())

		return nil
	},
}

var indexFlag int

func init() {
	rootCmd.AddCommand(dnsCmd)

	dnsCmd.AddCommand(listCmd)
	dnsCmd.AddCommand(addCmd)
	dnsCmd.AddCommand(deleteCmd)

	addCmd.Flags().IntVarP(&indexFlag, "index", "i", -1, "Index where the new DNS server will be inserted")
	deleteCmd.Flags().IntVarP(&indexFlag, "index", "i", -1, "Index of the DNS server to delete [required]")
	_ = deleteCmd.MarkFlagRequired("index")
}
