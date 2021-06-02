package start

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var port string

const long_description = `SoloDeploy CLI communicates with provider implementations over
gRPC. This command starts the gRPC listener and waits for the incoming messages at the port
provided by the --port flag.`

// NewCmd Go away linter
func NewCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "start",
		Short: "Starts the gRPC listener",
		Long:  long_description,
		Run: func(cmd *cobra.Command, args []string) {
			log.Print("Starting the gRPC server")
			log.Printf("Listening at localhost:%v", port)

			err := handler(port)
			if err != nil {
				log.Printf("Error starting the Solo GitHub GitProvider: %v", err)
				os.Exit(1)
			}
			os.Exit(0)
		},
	}
	c.Flags().StringVarP(&port, "port", "p", "", "The port number to use when opening the gRPC connection")
	c.MarkFlagRequired("port")

	return c
}
