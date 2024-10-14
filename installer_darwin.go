//go:build darwin
// +build darwin

package neuronos

import (
	_ "embed"
	"os"
	"path/filepath"
	"text/template"
)

var (
	//go:embed resources/start-on-boot.plist
	startOnBootPList string

	launchDir = filepath.Join(os.Getenv("HOME"), "Library", "LaunchAgents")
	execFile  = filepath.Join(os.Getenv("HOME"), "go/bin/application")

	AppData = struct {
		Name string
		Exec []string
	}{"NeuronOS", []string{execFile}}
)

func EnableStartOnBoot() error {
	// Check that file doesn't already exist
	path := filepath.Join(launchDir, "neuronos-demo.plist")
	if _, err := os.Stat(path); err != nil {
		return nil
	}

	// Load the plist template into memory
	tmpl := template.Must(template.New("").Parse(startOnBootPList))
	if err := os.MkdirAll(launchDir, 0777); err != nil {
		return err
	}

	// Create a destination file for the plist
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write to the file using the plist template
	return tmpl.Execute(file, AppData)
}
