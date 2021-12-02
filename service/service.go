package service

import (
	"log"

	"github.com/urfave/cli/v2"
)

const (
	SERVICE_NAME = "myService"
	SERVICE_DESC = "Dies ist ein Testservice!"
)

func Install(c *cli.Context) error {
	log.Print("calling installation function for you")
	err := install()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func Uninstall(c *cli.Context) error {
	log.Print("calling uninstallation function for you")
	err := uninstall()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func Execute(c *cli.Context) error {
	log.Print("i am executing the app for you")
	runService(SERVICE_NAME, false)
	return nil
}
