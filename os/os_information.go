package os

import (
	"bytes"
	"os/exec"
)

var (
	cmd    *exec.Cmd
	stdout bytes.Buffer
	stderr bytes.Buffer
)

type OSInformer interface {
	GetOSVersion() (string, error)
	GetHostname() (string, error)
	GetUptime() (string, error)
	GetNumberPackages() (string, error)
	GetShellInformation() (string, error)
	GetResolution() (string, error)
	GetDesktopEnvironment() (string, error)
	GetTerminalInfo() (string, error)
	GetCPU() (string, error)
	GetGPU() (string, error)
	GetMemoryUsage() (string, error)
}

type OSInformation struct {
	Name               string
	OS                 string
	Host               string
	Uptime             string
	Packages           string
	Shell              string
	Resolution         string
	DesktopEnvironment string
	Terminal           string
	CPU                string
	GPU                string
	Memory             string
}

//ExecuteCommand executes the command with arguments as well reset the stderr and stdout
func ExecuteCommand(command string, args ...string) (string, error) {
	// Reset stdout and stderr if previous commands were run
	stdout.Reset()
	stderr.Reset()
	cmd = exec.Command(command, args...)
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	err := cmd.Run()

	return stdout.String(), err
}
