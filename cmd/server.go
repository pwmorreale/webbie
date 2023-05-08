/*
Copyright © 2023 Peter W. Morreale

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/pwmorreale/webbie.git/internal/server"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "The server command starts the webbie web server.",
	Long:  `The server command starts the webbie web server.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := server.Start(cfgFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
