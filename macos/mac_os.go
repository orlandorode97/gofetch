package macos

import (
	"fmt"
	"os"
	"strings"

	osinfo "github.com/OrlandoRomo/gofetch/os"
	"github.com/shirou/gopsutil/mem"
)

type MacOS struct {
	osinfo.OSInformation
}

//GetInfo parse all the OSInformation fields
// TODO request OS information with go concurrency since it's taken around 2 o more to get simple information
func GetInfo() *MacOS {
	mac := MacOS{}

	if name, err := mac.GetName(); err == nil {
		name = strings.Replace(name, "\n", "", -1)
		mac.Name = name
	}

	if info, err := mac.GetOSVersion(); err == nil {
		info = strings.Replace(info, "\n", "", -1)
		mac.OS = info
	}

	if host, err := mac.GetHostname(); err == nil {
		mac.Host = host
	}

	if uptime, err := mac.GetUptime(); err == nil {
		uptime = strings.Replace(uptime, "\r\n", "", -1)
		uptimes := strings.Split(uptime, " ")
		mac.Uptime = uptimes[4] // Refactor this
	}

	if packages, err := mac.GetNumberPackages(); err == nil {
		packages = strings.TrimSpace(packages)
		mac.Packages = packages
	}

	if shell, err := mac.GetShellInformation(); err == nil {
		shell = strings.Replace(shell, "\n", "", -1)
		mac.Shell = shell
	}

	if resolution, err := mac.GetResolution(); err == nil {
		resolutions := strings.Split(resolution, "Resolution: ")
		resolution = strings.TrimSpace(resolutions[1])
		mac.Resolution = resolution
	}

	if de, err := mac.GetDesktopEnvironment(); err == nil {
		mac.DesktopEnvironment = de
	}

	if terminal, err := mac.GetTerminalInfo(); err == nil {
		terminal = strings.TrimSpace(terminal)
		mac.Terminal = terminal
	}

	if cpuInfo, err := mac.GetCPU(); err == nil {
		cpu := strings.Split(cpuInfo, ": ")
		cpuInfo = strings.Replace(cpu[1], "\n\r", "", -1)
		cpuInfo = strings.TrimSpace(cpuInfo)
		mac.CPU = cpuInfo
	}

	if gpu, err := mac.GetGPU(); err == nil {
		gpus := strings.Split(gpu, "Chipset Model: ")
		gpu = strings.TrimSpace(gpus[1])
		mac.GPU = gpu
	}

	if memory, err := mac.GetMemoryUsage(); err == nil {
		mac.Memory = memory
	}

	return &mac

}

// GetName returns the current user name
func (mac *MacOS) GetName() (string, error) {
	return osinfo.ExecuteCommand("whoami")
}

// GetOSVersion returns the name of the current OS, version and kernel version
func (mac *MacOS) GetOSVersion() (string, error) {
	return osinfo.ExecuteCommand("uname", "-srm")
}

// GetHostname returns the hostname of the machine
func (mac *MacOS) GetHostname() (string, error) {
	return os.Hostname()
}

// GetUptime returns the up time of the current OS
func (mac *MacOS) GetUptime() (string, error) {
	return osinfo.ExecuteCommand("uptime")
}

// GetNumberPackages return the number of packages install by homebrew
func (mac *MacOS) GetNumberPackages() (string, error) {
	command := "brew list --formula | wc -l"
	return osinfo.ExecuteCommand("bash", "-c", command)
}

// GetShellInformation return the used shell and its version
func (mac *MacOS) GetShellInformation() (string, error) {
	return osinfo.ExecuteCommand(os.ExpandEnv("$SHELL"), "--version")
}

// GetResolution returns the resolution of thee current monitor
func (mac *MacOS) GetResolution() (string, error) {
	command := "system_profiler SPDisplaysDataType  | grep Resolution"
	return osinfo.ExecuteCommand("bash", "-c", command)
}

// GetDesktopEnvironment returns the resolution of thee current monitor
func (mac *MacOS) GetDesktopEnvironment() (string, error) {
	return "Aqua", nil
}

// GetTerminalInfo get the current terminal name
func (mac *MacOS) GetTerminalInfo() (string, error) {
	return osinfo.ExecuteCommand("echo", os.ExpandEnv("$TERM_PROGRAM"))
}

// GetCPU returns the name of th CPU
func (mac *MacOS) GetCPU() (string, error) {
	command := "sysctl -a | grep machdep.cpu.brand_string"
	return osinfo.ExecuteCommand("bash", "-c", command)
}

// GetGPU returns the name of the GPU
func (mac *MacOS) GetGPU() (string, error) {
	command := "system_profiler SPDisplaysDataType | grep 'Chipset Model'"
	return osinfo.ExecuteCommand("bash", "-c", command)
}

// GetMemoryUsage returns the memory usage of the computer
func (mac *MacOS) GetMemoryUsage() (string, error) {
	memStat, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}
	total := memStat.Total / (1024 * 1024)
	used := memStat.Used / (1024 * 1024)
	return fmt.Sprintf("%v MB / %v MB", used, total), nil
}
