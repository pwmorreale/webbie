/*
Copyright © 2023 Peter W. Morreale

*/
package cmd

import (
	"os"

	"github.com/pwmorreale/webbie.git/internal/config"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "webbie",
	Short: "Webbie is a simple webserver that does nothing.",
	Long: `Webbie is a simple web server that literally does nothing.  It
serves a single web page that merely returns the headers (as JSON) 
of the invocation.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	defaultFile := config.ViperConfigPath + "/" + config.ViperConfigFileName

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is "+defaultFile+")")
}
