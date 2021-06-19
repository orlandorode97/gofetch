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
	currentOS := new(OSInformation)
	if info, err := getOSVersion(); err == nil {
		info = strings.Replace(info, "\n", "", -1)
		currentOS.Name = info
	}

	if host, err := os.Hostname(); err == nil {
		currentOS.Host = host
	}

	if uptime, err := getUptime(); err == nil {
		uptime = strings.Replace(uptime, "\r\n", "", -1)
		uptimes := strings.Split(uptime, " ")
		currentOS.Uptime = uptimes[4] // Refactor this
	}

	if packages, err := getNumberPackages(); err == nil {
		packages = strings.TrimSpace(packages)
		currentOS.Packages = packages
	}

	if shell, err := getShellInformation(); err == nil {
		shell = strings.Replace(shell, "\n", "", -1)
		currentOS.Shell = shell
	}

	if resolution, err := getResolution(); err == nil {
		resolutions := strings.Split(resolution, "Resolution: ")
		resolution = strings.TrimSpace(resolutions[1])
		currentOS.Resolution = resolution
	}

	if terminal, err := getTerminalInfo(); err == nil {
		terminal = strings.TrimSpace(terminal)
		currentOS.Terminal = terminal
	}

	if cpuInfo, err := getCPU(); err == nil {
		cpu := strings.Split(cpuInfo, ": ")
		cpuInfo = strings.Replace(cpu[1], "\n\r", "", -1)
		cpuInfo = strings.TrimSpace(cpuInfo)
		currentOS.CPU = cpuInfo
	}

	if gpu, err := getGPU(); err == nil {
		gpus := strings.Split(gpu, "Chipset Model: ")
		gpu = strings.TrimSpace(gpus[1])
		currentOS.GPU = gpu
	}

	if memory, err := getMemoryUsage(); err == nil {
		currentOS.Memory = memory
	}

	return currentOS

}

// getInfo returns the name of the current OS, version and kernel version
func getOSVersion() (string, error) {
	executeCommand("uname", "-srm")

	err := cmd.Run()

	return stdout.String(), err
}

// getUptime returns the up time of the current OS
func getUptime() (string, error) {
	executeCommand("uptime")

	err := cmd.Run()

	return stdout.String(), err
}

// getNumberPackages return the number of packages install by homebrew
func getNumberPackages() (string, error) {
	command := "brew list --formula | wc -l"
	executeCommand("bash", "-c", command)

	err := cmd.Run()

	return stdout.String(), err
}

// getShellInformation return the used shell and its version
func getShellInformation() (string, error) {
	executeCommand(os.ExpandEnv("$SHELL"), "--version")

	err := cmd.Run()

	return stdout.String(), err
}

// getResolution returns the resolution of thee current monitor
func getResolution() (string, error) {
	command := "system_profiler SPDisplaysDataType  | grep Resolution"
	executeCommand("bash", "-c", command)

	err := cmd.Run()

	return stdout.String(), err
}

// getTerminalInfo get the current terminal name
func getTerminalInfo() (string, error) {
	executeCommand("echo", os.ExpandEnv("$TERM_PROGRAM"))

	err := cmd.Run()

	return stdout.String(), err
}

// getCPU returns the name of th CPU
func getCPU() (string, error) {
	command := "sysctl -a | grep machdep.cpu.brand_string"
	executeCommand("bash", "-c", command)

	err := cmd.Run()

	return stdout.String(), err
}

// getGPU returns the name of the GPU
func getGPU() (string, error) {
	command := "system_profiler SPDisplaysDataType | grep 'Chipset Model'"
	executeCommand("bash", "-c", command)

	err := cmd.Run()

	return stdout.String(), err
}

// getMemoryUsage returns the memory usage of the computer
func getMemoryUsage() (string, error) {
	memStat, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}
	total := memStat.Total / (1024 * 1024)
	used := memStat.Used / (1024 * 1024)
	return fmt.Sprintf("%v MB / %v MB", used, total), nil
}

//executeCommand executes the command with arguments as well reset the stderr and stdout
func executeCommand(command string, args ...string) {
	// Reset stdout and stderr if previous commands were run
	stdout.Reset()
	stderr.Reset()
	cmd = exec.Command(command, args...)
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
}
