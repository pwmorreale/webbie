package server

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
	"github.com/spf13/viper"

	"github.com/pwmorreale/webbie.git/internal/config"
)

// Start is the entry point for the account gateway service
func Start(cfgFile string) error {

	err := config.InitConfig(cfgFile)
	if err != nil {
		return err
	}

	/*
		re := echo.New()
		re.HideBanner = true

		// Set routes...
		re.GET("/foo", func(c echo.Context) error {
			return Foo(c)
		})

	*/
	port := viper.GetString(config.ServerPort)

	fmt.Fprintln(os.Stdout, "port used is: ", port)
	return nil

	//return re.Start("80")
}

func Foo(c echo.Context) error {

	return nil
}
