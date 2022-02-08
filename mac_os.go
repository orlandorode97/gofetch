package gofetch

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/OrlandoRomo/gofetch/command"
	"github.com/shirou/gopsutil/mem"
)

var (
	regexUptime *regexp.Regexp
)

func init() {
	regexUptime = regexp.MustCompile(`\=\s(.*.+?),`)
}

type macos struct{}

func NewMacOS() command.Informer {
	return &macos{}
}

// GetName returns the current user name
func (mac *macos) GetName() (string, error) {
	return command.Execute("whoami")
}

// GetOSVersion returns the name of the current OS, version and kernel version
func (mac *macos) GetOSVersion() (string, error) {
	return command.Execute("uname", "-srm")
}

// GetHostname returns the hostname of the machine
func (mac *macos) GetHostname() (string, error) {
	return os.Hostname()
}

// GetUptime returns the up time of the current OS
func (mac *macos) GetUptime() (string, error) {
	boot, err := command.Execute("sysctl", "-n", "kern.boottime")
	if err != nil {
		return "", err
	}
	matches := regexUptime.FindStringSubmatch(boot)
	if len(matches) != 0 && matches[1] == "" {
		return "", nil
	}
	now := `$(date +%s)`
	seconds := fmt.Sprintf("echo $((%s - %s))", now, matches[1])
	seconds, err = command.Execute("bash", "-c", seconds)
	if err != nil {
		return "", err
	}

	return ParseUptime(seconds)
}

// GetNumberPackages return the number of packages install by homebrew
func (mac *macos) GetNumberPackages() (string, error) {
	cmd := "ls /usr/local/Cellar/ | wc -l"
	packages, err := command.Execute("bash", "-c", cmd)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(packages), nil
}

// GetShellInformation return the used shell and its version
func (mac *macos) GetShellInformation() (string, error) {
	shell, err := command.Execute(os.ExpandEnv("$SHELL"), "--version")
	if err != nil {
		return "", err
	}
	return shell, nil
}

// GetResolution returns the resolution of thee current monitor
func (mac *macos) GetResolution() (string, error) {
	cmd := "system_profiler SPDisplaysDataType  | grep Resolution"
	resolution, err := command.Execute("bash", "-c", cmd)
	if err != nil {
		return "", err
	}
	resolutions := strings.Split(resolution, "Resolution: ")
	resolution = strings.TrimSpace(resolutions[1])
	return resolution, nil
}

// GetDesktopEnvironment returns the resolution of thee current monitor
func (mac *macos) GetDesktopEnvironment() (string, error) {
	return "Aqua", nil
}

// GetTerminalInfo get the current terminal name
func (mac *macos) GetTerminalInfo() (string, error) {
	terminal, err := command.Execute("echo", os.ExpandEnv("$TERM_PROGRAM"))
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(terminal), nil
}

// GetCPU returns the name of th CPU
func (mac *macos) GetCPU() (string, error) {
	cmd := "sysctl -a | grep machdep.cpu.brand_string"
	cpuInfo, err := command.Execute("bash", "-c", cmd)
	if err != nil {
		return "", err
	}
	cpu := strings.Split(cpuInfo, ": ")
	cpuInfo = strings.Replace(cpu[1], "\n\r", "", -1)
	cpuInfo = strings.TrimSpace(cpuInfo)
	return cpuInfo, nil
}

// GetGPU returns the name of the GPU
func (mac *macos) GetGPU() (string, error) {
	cmd := "system_profiler SPDisplaysDataType | grep 'Chipset Model'"
	gpu, err := command.Execute("bash", "-c", cmd)
	if err != nil {
		return "", err
	}
	gpus := strings.Split(gpu, "Chipset Model: ")
	gpu = strings.TrimSpace(gpus[1])
	return gpu, nil
}

// GetMemoryUsage returns the memory usage of the computer
func (mac *macos) GetMemoryUsage() (string, error) {
	memStat, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}
	total := memStat.Total / (1024 * 1024)
	used := memStat.Used / (1024 * 1024)
	return fmt.Sprintf("%v MB / %v MB", used, total), nil
}
