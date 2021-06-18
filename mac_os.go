package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
)

var (
	cmd    *exec.Cmd
	stdout bytes.Buffer
	stderr bytes.Buffer
)

//GetInfo parse all the OSInformation fields
func GetInfo() *OSInformation {
	currentOS := new(OSInformation)
	if info, err := getInfo(); err == nil {
		info = strings.Replace(info, "\n", "", -1)
		currentOS.Name = info
	}

	if host, err := os.Hostname(); err == nil {
		currentOS.Host = host
	}

	if uptime, err := getUptime(); err == nil {
		uptime = strings.Replace(uptime, "\r\n", "", -1)
		uptimes := strings.Split(uptime, " ")
		currentOS.Uptime = uptimes[4]
	}

	if packages, err := getNumberPackages(); err == nil {
		packages = strings.TrimSpace(packages)
		currentOS.Packages = packages
	}

	if shell, err := getShellInformation(); err == nil {
		currentOS.Shell = shell
	}

	return currentOS

}

//getInfo returns the name of the current OS, version and kernel version
func getInfo() (string, error) {
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

func getShellInformation() (string, error) {
	executeCommand(os.ExpandEnv("$SHELL"), "--version")

	err := cmd.Run()

	return stdout.String(), err

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
