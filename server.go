package main

import (
	"github.com/frc-2175/benkins/server"
	"github.com/spf13/cobra"
)

func main() {
	var basePath string
	var password string
	cmd := &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			server.Main(basePath, password)
		},
	}
	cmd.Flags().StringVar(&basePath, "basePath", "", "The path to serve all files from")
	cmd.Flags().StringVar(&password, "password", "", "The password used for client authentication")

	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
