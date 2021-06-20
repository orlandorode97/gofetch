package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/shirou/gopsutil/mem"
)

var (
	cmd    *exec.Cmd
	stdout bytes.Buffer
	stderr bytes.Buffer
)

//GetInfo parse all the OSInformation fields
// TODO request OS information with go concurrency since it's taken around 2 o more to get simple information
func GetInfo() *OSInformation {
	currentOS := OSInformation{}

	if info, err := GetOSVersion(); err == nil {
		info = strings.Replace(info, "\n", "", -1)
		currentOS.Name = info
	}

	if host, err := GetHostname(); err == nil {
		currentOS.Host = host
	}

	if uptime, err := GetUptime(); err == nil {
		uptime = strings.Replace(uptime, "\r\n", "", -1)
		uptimes := strings.Split(uptime, " ")
		currentOS.Uptime = uptimes[4] // Refactor this
	}

	if packages, err := GetNumberPackages(); err == nil {
		packages = strings.TrimSpace(packages)
		currentOS.Packages = packages
	}

	if shell, err := GetShellInformation(); err == nil {
		shell = strings.Replace(shell, "\n", "", -1)
		currentOS.Shell = shell
	}

	if resolution, err := GetResolution(); err == nil {
		resolutions := strings.Split(resolution, "Resolution: ")
		resolution = strings.TrimSpace(resolutions[1])
		currentOS.Resolution = resolution
	}

	if de, err := GetDesktopEnvironment(); err == nil {
		currentOS.DesktopEnvironment = de
	}

	if terminal, err := GetTerminalInfo(); err == nil {
		terminal = strings.TrimSpace(terminal)
		currentOS.Terminal = terminal
	}

	if cpuInfo, err := GetCPU(); err == nil {
		cpu := strings.Split(cpuInfo, ": ")
		cpuInfo = strings.Replace(cpu[1], "\n\r", "", -1)
		cpuInfo = strings.TrimSpace(cpuInfo)
		currentOS.CPU = cpuInfo
	}

	if gpu, err := GetGPU(); err == nil {
		gpus := strings.Split(gpu, "Chipset Model: ")
		gpu = strings.TrimSpace(gpus[1])
		currentOS.GPU = gpu
	}

	if memory, err := GetMemoryUsage(); err == nil {
		currentOS.Memory = memory
	}

	return &currentOS

}

// GetOSVersion returns the name of the current OS, version and kernel version
func GetOSVersion() (string, error) {
	return executeCommand("uname", "-srm")
}

// GetHostname returns the hostname of the machine
func GetHostname() (string, error) {
	return os.Hostname()
}

// GetUptime returns the up time of the current OS
func GetUptime() (string, error) {
	return executeCommand("uptime")
}

// GetNumberPackages return the number of packages install by homebrew
func GetNumberPackages() (string, error) {
	command := "brew list --formula | wc -l"
	return executeCommand("bash", "-c", command)
}

// GetShellInformation return the used shell and its version
func GetShellInformation() (string, error) {
	return executeCommand(os.ExpandEnv("$SHELL"), "--version")
}

// GetResolution returns the resolution of thee current monitor
func GetResolution() (string, error) {
	command := "system_profiler SPDisplaysDataType  | grep Resolution"
	return executeCommand("bash", "-c", command)
}

// GetDesktopEnvironment returns the resolution of thee current monitor
func GetDesktopEnvironment() (string, error) {
	return "Aqua", nil
}

// GetTerminalInfo get the current terminal name
func GetTerminalInfo() (string, error) {
	return executeCommand("echo", os.ExpandEnv("$TERM_PROGRAM"))
}

// GetCPU returns the name of th CPU
func GetCPU() (string, error) {
	command := "sysctl -a | grep machdep.cpu.brand_string"
	return executeCommand("bash", "-c", command)
}

// GetGPU returns the name of the GPU
func GetGPU() (string, error) {
	command := "system_profiler SPDisplaysDataType | grep 'Chipset Model'"
	return executeCommand("bash", "-c", command)
}

// GetMemoryUsage returns the memory usage of the computer
func GetMemoryUsage() (string, error) {
	memStat, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}
	total := memStat.Total / (1024 * 1024)
	used := memStat.Used / (1024 * 1024)
	return fmt.Sprintf("%v MB / %v MB", used, total), nil
}

//executeCommand executes the command with arguments as well reset the stderr and stdout
func executeCommand(command string, args ...string) (string, error) {
	// Reset stdout and stderr if previous commands were run
	stdout.Reset()
	stderr.Reset()
	cmd = exec.Command(command, args...)
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	err := cmd.Run()

	return stdout.String(), err
}
