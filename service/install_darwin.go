package service

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"text/template"
)

//go:embed scripts/launchdScript
var script string

type templateData struct {
	Label         string
	ExecutionPath string
}

func install() error {
	log.Print("reading embedded script...\n")

	log.Print("trying to get the working directory...\n")
	wd, err := os.Getwd()
	if err != nil {
		log.Print("fehler beim feststellen der working directory\n")
		return err
	}
	log.Printf("success. working directory is: %s\n", wd)
	ep := fmt.Sprintf("%s/TestApp", wd)
	log.Printf("successful. execution path therefore is: %s\n", ep)

	log.Print("initialising template data")
	td := templateData{
		Label:         "TestApp",
		ExecutionPath: ep,
	}

	log.Print("configuring configuration path to the property list file\n")
	configPath := fmt.Sprintf("/Library/LaunchDaemons/%s.plist", "TestApp")

	_, err = os.Stat(configPath)
	if err == nil {
		return fmt.Errorf("configuration already exists at: %s. please uninstall first. ", configPath)
	}

	log.Print("trying to open the property list file\n")
	f, err := os.Create(configPath)
	if err != nil {
		log.Print("fehler beim erstellen der datei\n")
		return err
	}
	log.Print("success.\n")

	log.Print("trying to create a new template\n")
	t := template.Must(template.New("launchdConfig").Parse(script))
	log.Print("success. trying to execute the template\n")
	err = t.Execute(f, td)
	if err != nil {
		log.Print("fehler beim ausführen des templates\n")
		return err
	}
	err = os.Chmod(configPath, 0755)
	if err != nil {
		log.Print("fehler beim ändern der berechtigungen\n")
		return err
	}
	log.Print("success. service should be installed\n")

	return nil
}

func uninstall() error {
	configPath := "/Library/LaunchDaemons/TestApp.plist"
	log.Print("versuche die property list datei zu entfernen")
	err := os.Remove(configPath)
	if err != nil {
		log.Print("fehler beim entfernen der property list datei")
		return err
	}
	log.Print("erfolgreich entfernt.")
	return nil
}
