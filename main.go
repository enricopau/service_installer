package main

import (
	"fmt"
	"os"

	"github.com/enricopau/service_installer/beep"
	"github.com/enricopau/service_installer/service"

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
			{
				Name:   "app",
				Usage:  "starts as an application",
				Action: beep.Beep,
			},
		},
	}
	if len(os.Args) <= 0 {
		os.Args = append(os.Args, "app")
		app.Run(os.Args)
	}
	app.Run(os.Args)

	// isWinSvc, err := svc.IsWindowsService()
	// if err != nil {
	// 	log.Print("kann ich nicht herausfinden :( srüüüüüüüüü")
	// }
	// if isWinSvc {
	// 	beep.Beep()
	// }

}
