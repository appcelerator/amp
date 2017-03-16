package main

import (
	"fmt"
	"strings"

	"github.com/appcelerator/amp/api/rpc/service"
	"github.com/appcelerator/amp/cmd/amp/cli"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	// TODO: add support for 'remove' alias
	serviceRmCmd = &cobra.Command{
		Use:     "rm",
		Short:   "Remove one or more services",
		Example: "sample-service",
		Aliases: []string{"del"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return serviceRm(AMP, args)
		},
	}
)

func init() {
	ServiceCmd.AddCommand(serviceRmCmd)
}

func serviceRm(amp *cli.AMP, args []string) error {
	if len(args) < 1 {
		mgr.Fatal("\"amp service rm\" requires at least 1 argument(s)")
	}

	client := service.NewServiceClient(amp.Conn)

	for _, ident := range args {
		req := &service.RemoveRequest{
			Ident: ident,
		}

		resp, err := client.Remove(context.Background(), req)
		if err != nil {
			pattern := "daemon: "
			errstr := fmt.Sprintf("%s", err)
			index := strings.LastIndex(errstr, pattern)
			if index > -1 {
				mgr.Fatal("error : %s", errstr[index+len(pattern):])
			} else {
				mgr.Fatal("error : %s", grpc.ErrorDesc(err))
			}
		} else {
			fmt.Println(resp.Ident)
		}
	}

	return nil
}
