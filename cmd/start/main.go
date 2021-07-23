package start

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	org  string
	pat  string
	port string
)

const longDescription = `SoloDeploy CLI communicates with provider implementations over
gRPC. This command starts the gRPC listener and waits for the incoming messages at the port
provided by the --port flag.`

// NewCmd Go away linter
func NewCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "start",
		Short: "Starts the gRPC listener",
		Long:  longDescription,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Starting the Solo GitHub Provider gRPC server")
			fmt.Printf("Listening at localhost:%v\n", port)

			err := handler(port, org, pat)
			if err != nil {
				fmt.Printf("Error starting the Solo GitHub GitProvider: %v\n", err)
				os.Exit(1)
			}
			os.Exit(0)
		},
	}
	c.Flags().StringVarP(&port, "port", "p", "", "The port number to use when opening the gRPC connection")
	c.Flags().StringVarP(&org, "organization", "o", "", "The github organisation to work with")
	c.Flags().StringVarP(&pat, "pat", "", "", "The github personal access token to authenticate with")
	c.MarkFlagRequired("port")
	c.MarkFlagRequired("organization")
	c.MarkFlagRequired("pat")

	return c
}
