package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Constants for various configuration parameters.
const (

	// Viper configuration.
	ViperConfigPath     = "/etc/webbie"
	ViperConfigFileName = "conf"
	ViperConfigFileType = "yaml"
	ViperConfigEnv      = "WEBBIE_CONFIG_PATH"

	// REST Server configuration
	ServerPort = "server.port"
)

// initConfig reads in config file and ENV variables if set.
// Note that webbie uses the global viper instance.
func InitConfig(cfgFile string) error {

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Add the path specified by the env var.
		path, ok := os.LookupEnv(ViperConfigEnv)
		if ok {
			viper.AddConfigPath(path)
		} else {
			viper.AddConfigPath(ViperConfigPath)
			viper.SetConfigType(ViperConfigFileType)
			viper.SetConfigName(ViperConfigFileName)
		}
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	return err
}
