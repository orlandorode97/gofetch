package linux

import (
	"fmt"
	"os"

	osinfo "github.com/OrlandoRomo/gofetch/os"
	"github.com/shirou/gopsutil/mem"
)

type PackageManager string
type Command string

var distrosPackages map[PackageManager]Command

const NetPackage = `which {xbps-install,apk,apt,pacman,nix,yum,rpm,dpkg,emerge} 2>/dev/null | grep -v "not found" | awk -F/ 'NR==1{print $NF}')"`

func init() {
	distrosPackages = map[PackageManager]Command{
		"xbps-install": "xbps-query -l | wc -l",
		"apk":          "apk search | wc -l",
		"apt":          "apt list --installed 2>/dev/null | wc -l",
		"pacman":       "pacman -Q | wc -l",
		"nix":          `nix-env -qa --installed "*" | wc -l`,
		"yum":          "yum list installed | wc -l",
		"rpm":          "rpm -qa | wc -l",
		"emerge":       "qlist -I | wc -l",
	}
}

type Linux struct {
	osinfo.OSInformation
}

func GetInfo() *Linux {
	linux := Linux{}
	return &linux
}

// GetName returns the current user name
func (l *Linux) GetName() (string, error) {
	return osinfo.ExecuteCommand("whoami")
}

// GetOSVersion returns the name of the current OS, version and kernel version
func (l *Linux) GetOSVersion() (string, error) {
	return osinfo.ExecuteCommand("uname", "-srm")
}

// GetHostname returns the hostname of the machine
func (l *Linux) GetHostname() (string, error) {
	return os.Hostname()
}

// GetUptime returns the up time of the current OS
func (l *Linux) GetUptime() (string, error) {
	return osinfo.ExecuteCommand("uptime")
}

// GetNumberPackages return the number of packages install by homebrew
func (l *Linux) GetNumberPackages() (string, error) {
	packageManager, err := osinfo.ExecuteCommand("bash", "-c", NetPackage)
	if err != nil {
		return "", err
	}

	name, ok := distrosPackages[PackageManager(packageManager)]

	if !ok {
		return "Unknown", nil
	}

	return osinfo.ExecuteCommand("bash", "-c", string(name))
}

// GetShellInformation return the used shell and its version
func (l *Linux) GetShellInformation() (string, error) {
	return osinfo.ExecuteCommand(os.ExpandEnv("$SHELL"), "--version")
}

// GetResolution returns the resolution of thee current monitor
func (l *Linux) GetResolution() (string, error) {
	command := "xdpyinfo | grep 'dimensions:'"
	return osinfo.ExecuteCommand("bash", "-c", command)
}

// GetDesktopEnvironment returns the resolution of the current monitor
func (l *Linux) GetDesktopEnvironment() (string, error) {
	return "Aqua", nil
}

// GetTerminalInfo get the current terminal name
func (l *Linux) GetTerminalInfo() (string, error) {
	return osinfo.ExecuteCommand("echo", os.ExpandEnv("$TERM_PROGRAM"))
}

// GetCPU returns the name of th CPU
func (l *Linux) GetCPU() (string, error) {
	command := "sysctl -a | grep machdep.cpu.brand_string"
	return osinfo.ExecuteCommand("bash", "-c", command)
}

// GetGPU returns the name of the GPU
func (l *Linux) GetGPU() (string, error) {
	command := "system_profiler SPDisplaysDataType | grep 'Chipset Model'"
	return osinfo.ExecuteCommand("bash", "-c", command)
}

// GetMemoryUsage returns the memory usage of the computer
func (l *Linux) GetMemoryUsage() (string, error) {
	memStat, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}
	total := memStat.Total / (1024 * 1024)
	used := memStat.Used / (1024 * 1024)
	return fmt.Sprintf("%v MB / %v MB", used, total), nil
}
