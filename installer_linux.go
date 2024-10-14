//go:build linux
// +build linux

package neuronos

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

var (
	//go:embed resources/start-on-boot.service
	startOnBootService string

	systemdDir = "/etc/systemd/system"
	execFile   = filepath.Join(os.Getenv("HOME"), "go/bin/application")

	AppData = struct {
		Name string
		Exec []string
	}{"NeuronOS", []string{execFile}}
)

func EnableStartOnBoot() error {
	// Path to the systemd service file
	path := filepath.Join(systemdDir, "neuronos-demo.service")
	if _, err := os.Stat(path); err == nil {
		// If the service file exists, return without doing anything
		fmt.Println("Service file already exists")
		return nil
	}

	// Load the systemd service template into memory
	tmpl := template.Must(template.New("").Parse(startOnBootService))
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	// Write to the file using the systemd service template
	if err := tmpl.Execute(file, AppData); err != nil {
		return err
	}

	// Reload the systemd manager configuration
	if err := bash("systemctl", "daemon-reload"); err != nil {
		return err
	}

	// Enable the service so it starts on boot
	if err := bash("systemctl", "enable", "neuronos-demo.service"); err != nil {
		return err
	}

	// Start the service immediately
	return bash("systemctl", "start", "neuronos-demo.service")
}

// Helper function to run system commands
func bash(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
