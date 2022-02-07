package gofetch

import (
	"fmt"
	"os"

	"github.com/OrlandoRomo/gofetch/command"
	"github.com/shirou/gopsutil/mem"
)

type windows struct{}

// NewWin returns an instance of the interface OSInformer
func NewWin() command.Informer {
	return &windows{}
}

// GetName returns the current user name
func (mac *windows) GetName() (string, error) {
	return command.Execute("whoami")
}

// GetOSVersion returns the name of the current OS, version and kernel version
func (mac *windows) GetOSVersion() (string, error) {
	return command.Execute("uname", "-srm")
}

// GetHostname returns the hostname of the machine
func (mac *windows) GetHostname() (string, error) {
	return os.Hostname()
}

// GetUptime returns the up time of the current OS
func (mac *windows) GetUptime() (string, error) {
	return command.Execute("uptime")
}

// GetNumberPackages return the number of packages install by homebrew
func (mac *windows) GetNumberPackages() (string, error) {
	cmd := "brew list --formula | wc -l"
	return command.Execute("bash", "-c", cmd)
}

// GetShellInformation return the used shell and its version
func (mac *windows) GetShellInformation() (string, error) {
	return command.Execute(os.ExpandEnv("$SHELL"), "--version")
}

// GetResolution returns the resolution of thee current monitor
func (mac *windows) GetResolution() (string, error) {
	cmd := "system_profiler SPDisplaysDataType  | grep Resolution"
	return command.Execute("bash", "-c", cmd)
}

// GetDesktopEnvironment returns the resolution of thee current monitor
func (mac *windows) GetDesktopEnvironment() (string, error) {
	return "Aqua", nil
}

// GetTerminalInfo get the current terminal name
func (mac *windows) GetTerminalInfo() (string, error) {
	return command.Execute("echo", os.ExpandEnv("$TERM_PROGRAM"))
}

// GetCPU returns the name of th CPU
func (mac *windows) GetCPU() (string, error) {
	cmd := "sysctl -a | grep machdep.cpu.brand_string"
	return command.Execute("bash", "-c", cmd)
}

// GetGPU returns the name of the GPU
func (mac *windows) GetGPU() (string, error) {
	cmd := "system_profiler SPDisplaysDataType | grep 'Chipset Model'"
	return command.Execute("bash", "-c", cmd)
}

// GetMemoryUsage returns the memory usage of the computer
func (mac *windows) GetMemoryUsage() (string, error) {
	memStat, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}
	total := memStat.Total / (1024 * 1024)
	used := memStat.Used / (1024 * 1024)
	return fmt.Sprintf("%v MB / %v MB", used, total), nil
}
