package windows

import (
	"fmt"
	"os"

	"github.com/OrlandoRomo/gofetch/command"
	"github.com/shirou/gopsutil/mem"
)

type Windows struct{}

func New() *Windows {
	return &Windows{}
}

// GetName returns the current user name
func (mac *Windows) GetName() (string, error) {
	return command.ExecuteCommand("whoami")
}

// GetOSVersion returns the name of the current OS, version and kernel version
func (mac *Windows) GetOSVersion() (string, error) {
	return command.ExecuteCommand("uname", "-srm")
}

// GetHostname returns the hostname of the machine
func (mac *Windows) GetHostname() (string, error) {
	return os.Hostname()
}

// GetUptime returns the up time of the current OS
func (mac *Windows) GetUptime() (string, error) {
	return command.ExecuteCommand("uptime")
}

// GetNumberPackages return the number of packages install by homebrew
func (mac *Windows) GetNumberPackages() (string, error) {
	cmd := "brew list --formula | wc -l"
	return command.ExecuteCommand("bash", "-c", cmd)
}

// GetShellInformation return the used shell and its version
func (mac *Windows) GetShellInformation() (string, error) {
	return command.ExecuteCommand(os.ExpandEnv("$SHELL"), "--version")
}

// GetResolution returns the resolution of thee current monitor
func (mac *Windows) GetResolution() (string, error) {
	cmd := "system_profiler SPDisplaysDataType  | grep Resolution"
	return command.ExecuteCommand("bash", "-c", cmd)
}

// GetDesktopEnvironment returns the resolution of thee current monitor
func (mac *Windows) GetDesktopEnvironment() (string, error) {
	return "Aqua", nil
}

// GetTerminalInfo get the current terminal name
func (mac *Windows) GetTerminalInfo() (string, error) {
	return command.ExecuteCommand("echo", os.ExpandEnv("$TERM_PROGRAM"))
}

// GetCPU returns the name of th CPU
func (mac *Windows) GetCPU() (string, error) {
	cmd := "sysctl -a | grep machdep.cpu.brand_string"
	return command.ExecuteCommand("bash", "-c", cmd)
}

// GetGPU returns the name of the GPU
func (mac *Windows) GetGPU() (string, error) {
	cmd := "system_profiler SPDisplaysDataType | grep 'Chipset Model'"
	return command.ExecuteCommand("bash", "-c", cmd)
}

// GetMemoryUsage returns the memory usage of the computer
func (mac *Windows) GetMemoryUsage() (string, error) {
	memStat, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}
	total := memStat.Total / (1024 * 1024)
	used := memStat.Used / (1024 * 1024)
	return fmt.Sprintf("%v MB / %v MB", used, total), nil
}
