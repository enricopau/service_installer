package service

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
)

type winsvc struct {
	app *cli.App
}

func (ws *winsvc) Execute(args []string, r <-chan svc.ChangeRequest, changes chan<- svc.Status) (ssec bool, errno uint32) {
	//Select which service commands can be accepted to control the service.
	//This is a minimal set and needs to be expanded if functionality will be added.
	const cmdAcc = svc.AcceptStop | svc.AcceptShutdown
	//Send a new service status in the changes channel for StartPending
	changes <- svc.Status{State: svc.StartPending}

	//start the service
	errCh := make(chan error)
	go func() {
		errCh <- ws.app.Run(os.Args)
	}()
	// err := ws.StartService()
	// if err != nil {
	// 	return
	// }
	//Send a new service status in the changes channel for Running
	changes <- svc.Status{State: svc.Running, Accepts: cmdAcc}
	var err error
loop:
	for {
		select {
		case c := <-r:
			switch c.Cmd {
			case svc.Interrogate:
				changes <- c.CurrentStatus
			case svc.Stop:
				changes <- svc.Status{State: svc.StopPending}
				err = ws.StopService()
				if err != nil {
					return
				}
				break loop
			case svc.Shutdown:
				changes <- svc.Status{State: svc.StopPending}
				err = ws.StopService()
				if err != nil {
					return
				}
				break loop
			default:
				continue loop
			}
		case err = <-errCh:
			if err != nil {
				log.Print(1, fmt.Sprintf("terminated with error %v", err))
				ssec = true
				errno = 1
			} else {
				log.Print(1, "terminated without error")
				errno = 0
			}
			return

		}

	}
	return
}

func runService(name string, isDebug bool) {
	var err error

	run := svc.Run
	err = run(name, &winsvc{})
	if err != nil {
		return
	}
}

//StartService can start a service via another interface.
//This should be analogous to what the windows service application does.
func StartService() error {
	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()
	s, err := m.OpenService(SERVICE_NAME)
	if err != nil {
		return err
	}
	defer s.Close()
	return s.Start()
}

func (ws *winsvc) StopService() error {
	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()
	s, err := m.OpenService(SERVICE_NAME)
	if err != nil {
		return err
	}
	defer s.Close()
	status, err := s.Control(svc.Stop)
	if err != nil {
		return err
	}
	timeout := time.Now().Add(10 * time.Second)
	for status.State != svc.Stopped {
		if timeout.Before(time.Now()) {
			return fmt.Errorf("Warte, dass der Service stoppt:%s", err)
		}
		time.Sleep(300 * time.Millisecond)
		status, err = s.Query()
		if err != nil {
			return fmt.Errorf("Konnte den Servicestatus nicht ermitteln: %s", err)
		}
	}
	return nil

}
