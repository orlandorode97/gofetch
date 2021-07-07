package linux

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	osinfo "github.com/OrlandoRomo/gofetch/os"
	"github.com/shirou/gopsutil/mem"
)

type PackageManager string
type Command string

var (
	distrosPackages map[PackageManager]Command
	regexGPU        *regexp.Regexp
	regexPackages   *regexp.Regexp
)

// command to found the binary file of the current package manager
const NetPackage = `which {xbps-install,apk,apt,pacman,nix,yum,rpm,emerge} 2>/dev/null | grep -v "not found"`

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
	regexGPU = regexp.MustCompile(`(Intel|Advanced|NVIDIA|MCST|Virtual Box)([^\(|\(|\\]+)`)
	regexPackages = regexp.MustCompile(`[^/]*$`)
}

type Linux struct {
	osinfo.OSInformation
}

func GetInfo() *Linux {
	linux := Linux{}
	if name, err := linux.GetName(); err == nil {
		linux.Name = name
	}

	if info, err := linux.GetOSVersion(); err == nil {
		linux.OS = info
	}

	if host, err := linux.GetHostname(); err == nil {
		linux.Host = host
	}

	if uptime, err := linux.GetUptime(); err == nil {
		uptime = strings.Replace(uptime, "\r\n", "", -1)
		uptimes := strings.Split(uptime, " ")
		linux.Uptime = uptimes[4] // Refactor this
	}

	if packages, err := linux.GetNumberPackages(); err == nil {
		linux.Packages = packages
	}

	if shell, err := linux.GetShellInformation(); err == nil {
		linux.Shell = shell
	}

	if resolution, err := linux.GetResolution(); err == nil {
		resolutions := strings.Split(resolution, "dimensions: ")
		resolution = strings.TrimSpace(resolutions[1])
		linux.Resolution = resolution
	}

	if de, err := linux.GetDesktopEnvironment(); err == nil {
		linux.DesktopEnvironment = de
	}

	if terminal, err := linux.GetTerminalInfo(); err == nil {
		terminal = strings.TrimSpace(terminal)
		linux.Terminal = terminal
	}

	if cpuInfo, err := linux.GetCPU(); err == nil {
		cpu := strings.Split(cpuInfo, ": ")
		cpuInfo = strings.Replace(cpu[1], "\n\r", "", -1)
		cpuInfo = strings.TrimSpace(cpuInfo)
		linux.CPU = cpuInfo
	}

	if gpu, err := linux.GetGPU(); err == nil {
		if regexGPU.MatchString(gpu) {
			gpu = regexGPU.FindString(gpu)
			linux.GPU = gpu
		}
	}

	if memory, err := linux.GetMemoryUsage(); err == nil {
		linux.Memory = memory
	}

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

// GetHostname returns the hostname of the linuxhine
func (l *Linux) GetHostname() (string, error) {
	return os.Hostname()
}

// GetUptime returns the up time of the current OS
func (l *Linux) GetUptime() (string, error) {
	return osinfo.ExecuteCommand("uptime")
}

// GetNumberPackages return the number of packages install by homebrew
func (l *Linux) GetNumberPackages() (string, error) {
	packageManager, err := osinfo.ExecuteCommand(`bash`, `-c`, NetPackage)
	if err != nil {
		return "", err
	}

	if regexPackages.MatchString(packageManager) {
		packageManager = regexPackages.FindString(packageManager)
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
	command := "lscpu | grep 'Model name:'"
	return osinfo.ExecuteCommand("bash", "-c", command)
}

// GetGPU returns the name of the GPU
func (l *Linux) GetGPU() (string, error) {
	command := "lspci -v | grep 'VGA\\|Display\\|3D'"
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