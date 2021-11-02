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
const NetPackage = `which {xbps-install,apk,dpkg,pacman,nix,yum,rpm,emerge} 2>/dev/null | grep -v "not found"`

func init() {
	distrosPackages = map[PackageManager]Command{
		"xbps-install": "xbps-query -l | wc -l",
		"apk":          "apk search | wc -l",
		"dpkg":         "dpkg-query -f '.\n' -W | wc -l",
		"pacman":       "pacman -Q | wc -l",
		"nix":          `nix-env -qa --installed "*" | wc -l`,
		"yum":          "yum list installed | wc -l",
		"rpm":          "rpm -qa | wc -l",
		"emerge":       "qlist -I | wc -l",
	}
	regexGPU = regexp.MustCompile(`(Intel|Advanced|NVIDIA|MCST|Virtual Box)([^\(|\(|\\]+)`)
	regexPackages = regexp.MustCompile(`[^/]*$`)
}

type Linux struct{}

func NewLinux() *Linux {
	return &Linux{}
}

// GetName returns the current user name
func (l *Linux) GetName() (string, error) {
	return osinfo.ExecuteCommand("whoami")
}

// GetOSVersion returns the name of the current OS, version and kernel version
func (l *Linux) GetOSVersion() (string, error) {
	return osinfo.ExecuteCommand("uname", "-srm")
}

// GetHostname returns the hostname of the linux distro
func (l *Linux) GetHostname() (string, error) {
	return os.Hostname()
}

// GetUptime returns the up time of the current OS
func (l *Linux) GetUptime() (string, error) {
	uptime, err := osinfo.ExecuteCommand("uptime")
	if err != nil {
		return "", err
	}
	uptime = strings.Replace(uptime, "\r\n", "", -1)
	uptimes := strings.Split(uptime, " ")
	return uptimes[4], nil
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
	command := fmt.Sprintf("echo %s | awk -F'/' '{print $NF}'", os.ExpandEnv("$SHELL"))
	shell, err := osinfo.ExecuteCommand("bash", "-c", command)
	if err != nil {
		return "", err
	}
	return shell, nil
}

// GetResolution returns the resolution of thee current monitor
func (l *Linux) GetResolution() (string, error) {
	command := "xdpyinfo | grep 'dimensions:'"
	resolution, err := osinfo.ExecuteCommand("bash", "-c", command)
	if err != nil {
		return "", err
	}
	resolutions := strings.Split(resolution, "dimensions: ")
	resolution = strings.TrimSpace(resolutions[1])
	return resolution, nil

}

// GetDesktopEnvironment returns the resolution of the current monitor
func (l *Linux) GetDesktopEnvironment() (string, error) {
	return "Aqua", nil
}

// GetTerminalInfo get the current terminal name
func (l *Linux) GetTerminalInfo() (string, error) {
	terminal, err := osinfo.ExecuteCommand("echo", os.ExpandEnv("$TERM"))
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(terminal), nil
}

// GetCPU returns the name of th CPU
func (l *Linux) GetCPU() (string, error) {
	command := "lscpu | grep 'Model name:'"
	cpuInfo, err := osinfo.ExecuteCommand("bash", "-c", command)
	if err != nil {
		return "", err
	}
	cpu := strings.Split(cpuInfo, ": ")
	cpuInfo = strings.Replace(cpu[1], "\n\r", "", -1)
	cpuInfo = strings.TrimSpace(cpuInfo)
	return cpuInfo, nil
}

// GetGPU returns the name of the GPU
func (l *Linux) GetGPU() (string, error) {
	command := "lspci -v | grep 'VGA\\|Display\\|3D'"
	gpu, err := osinfo.ExecuteCommand("bash", "-c", command)
	if err != nil {
		return "", err
	}
	if regexGPU.MatchString(gpu) {
		gpu = regexGPU.FindString(gpu)
	}
	return gpu, nil
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
