package service

import (
	"log"

	"github.com/urfave/cli/v2"
)

const (
	SERVICE_NAME = "myService"
	SERVICE_DESC = "Dies ist ein Testservice!"
)

//Install will be called and will call the os-specific install function.
func Install(c *cli.Context) error {
	log.Print("calling installation function for you")
	err := install()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

//Uninstall will be called and will call the os-specific uninstall function.
func Uninstall(c *cli.Context) error {
	log.Print("calling uninstallation function for you")
	err := uninstall()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
