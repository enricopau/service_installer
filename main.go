package main

import (
	"fmt"
	"log"
	"os"

	"github.com/enricopau/service_installer/beep"
	"github.com/enricopau/service_installer/service"
	"golang.org/x/sys/windows/svc"

	"github.com/urfave/cli/v2"
)

func main() {
	fmt.Println("Hello, user.")

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "service",
				Usage: "controls the service actions, use 'install' or 'uninstall' as second parameters",
				Subcommands: []*cli.Command{
					{
						Name:   "install",
						Usage:  "installs a service on your operating system",
						Action: service.Install,
					},
					{
						Name:   "uninstall",
						Usage:  "uninstalls a service from your operating system",
						Action: service.Uninstall,
					},
				},
			},
		},
	}
	app.Run(os.Args)

	isInteractive, err := svc.IsAnInteractiveSession()
	if err != nil {
		log.Print("kann ich nicht herausfinden :( srüüüüüüüüü")
	}
	if isInteractive {
		beep.Beep()
	}

}
