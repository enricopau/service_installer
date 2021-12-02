package service

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows/svc/eventlog"

	"golang.org/x/sys/windows/svc/mgr"
)

func install() error {
	log.Print("starting to install the service on your operating system...")

	ep, err := execPath()
	if err != nil {
		return err
	}

	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()
	_, err = m.OpenService(SERVICE_NAME)
	if err == nil {
		return fmt.Errorf("service %s already exists", SERVICE_NAME)
	}
	s, err := m.CreateService(SERVICE_NAME, ep, mgr.Config{
		Description:      SERVICE_DESC,
		StartType:        mgr.StartAutomatic,
		ServiceType:      0,                            //0 value will default to Win32OwnProcess; see: https://docs.microsoft.com/de-de/dotnet/api/system.serviceprocess.servicetype?view=dotnet-plat-ext-6.0
		ServiceStartName: "NT AUTHORITY\\LocalService", //only set this if you don't want your service to run under the local system account
	})
	if err != nil {
		return err
	}
	defer s.Close()
	err = eventlog.InstallAsEventCreate(SERVICE_NAME, eventlog.Error|eventlog.Warning|eventlog.Info)
	if err != nil {
		s.Delete()
		return err
	}
	log.Print("successfully installed the service on your operating system...")
	return nil
}

func uninstall() error {
	log.Print("starting to remove the service from your operating system...")
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
	err = s.Delete()
	if err != nil {
		return err
	}
	err = eventlog.Remove(SERVICE_NAME)
	if err != nil {
		return err
	}
	log.Print("successfully removed the service on your operating system...")
	return nil
}

// execPath can be used if you want to use the application in which this service package is embedded as the service itself.
// Otherwise user the path to the executable.
func execPath() (string, error) {
	progName := os.Args[0]
	progPath, err := filepath.Abs(progName)
	if err != nil {
		return "", err
	}
	fi, err := os.Stat(progPath)
	if err == nil {
		if !fi.Mode().IsDir() {
			return progPath, nil
		}
		err = fmt.Errorf("%s is directory", progPath)
	}
	if filepath.Ext(progPath) == "" {
		progPath += ".exe"
		fi, err := os.Stat(progPath)
		if err == nil {
			if !fi.Mode().IsDir() {
				return progPath, nil
			}
			err = fmt.Errorf("%s is directory", progPath)
			return "", err
		}
	}
	return "", err
}
