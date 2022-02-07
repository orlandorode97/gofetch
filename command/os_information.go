package command

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"

	"github.com/fatih/color"
)

var (
	cmd *exec.Cmd
)

var (
	Red     *color.Color
	Green   *color.Color
	Cyan    *color.Color
	Yellow  *color.Color
	Blue    *color.Color
	Magenta *color.Color
	White   *color.Color
)

func init() {
	Red = color.New(color.FgRed, color.Bold)
	Green = color.New(color.FgGreen, color.Bold)
	Cyan = color.New(color.FgCyan, color.Bold)
	Yellow = color.New(color.FgYellow, color.Bold)
	Blue = color.New(color.FgBlue, color.Bold)
	Magenta = color.New(color.FgHiMagenta, color.Bold)
	White = color.New(color.FgWhite, color.Bold)

}

type Versioner interface {
	GetOSVersion() (string, error)
}

type Namer interface {
	GetName() (string, error)
}

type Hoster interface {
	GetHostname() (string, error)
}

type Timer interface {
	GetUptime() (string, error)
}

type Packager interface {
	GetNumberPackages() (string, error)
}

type Sheller interface {
	GetShellInformation() (string, error)
}

type Resolusioner interface {
	GetResolution() (string, error)
}

type Environment interface {
	GetDesktopEnvironment() (string, error)
}

type Terminal interface {
	GetTerminalInfo() (string, error)
}

type CPU interface {
	GetCPU() (string, error)
}

type GPU interface {
	GetGPU() (string, error)
}

type Usager interface {
	GetMemoryUsage() (string, error)
}

type Informer interface {
	Versioner
	Namer
	Hoster
	Timer
	Packager
	Sheller
	Resolusioner
	Environment
	Terminal
	CPU
	GPU
	Usager
}

//ExecuteCommand executes the command with arguments as well reset the stderr and stdout
func Execute(command string, args ...string) (string, error) {
	var bufOut BufferOut
	var bufErr BufferErr
	// Reset stdout and stderr if previous commands were run
	bufOut.Reset()
	bufErr.Reset()
	cmd = exec.Command(command, args...)
	cmd.Stderr = &bufErr.stderr
	cmd.Stdout = &bufOut.stdout

	err := cmd.Run()

	return strings.TrimSuffix(bufOut.stdout.String(), "\n"), err
}

func Fetch(in Informer) {

	waitGroup := sync.WaitGroup{}

	waitGroup.Add(9)
	if name, err := in.GetName(); err == nil {
		fmt.Printf("%s", Red.Sprint(name))
	}
	if host, err := in.GetHostname(); err == nil {
		fmt.Printf("@%s\n", Red.Sprint(host))
	}
	fmt.Printf("%s %s %s %s %s\n\n", Red.Sprint("X"), Green.Sprint("────"), Yellow.Sprint("X"), Green.Sprint("────"), Blue.Sprint("X"))

	go func() {
		if uptime, err := in.GetUptime(); err == nil {
			fmt.Printf("%s %s %s\n", Cyan.Sprint("uptime"), "~", uptime)
		}
		waitGroup.Done()
	}()
	go func() {
		if numPackages, err := in.GetNumberPackages(); err == nil {
			fmt.Printf("%s %s %s\n", Blue.Sprint("packages"), "~", numPackages)
		}
		waitGroup.Done()
	}()
	go func() {
		if shell, err := in.GetShellInformation(); err == nil {
			fmt.Printf("%s %s %s\n", Yellow.Sprint("shell"), "~", shell)
		}
		waitGroup.Done()
	}()
	go func() {
		if resolution, err := in.GetResolution(); err == nil {
			fmt.Printf("%s %s %s\n", Red.Sprint("resolution"), "~", resolution)
		}
		waitGroup.Done()
	}()
	go func() {
		if deskEnv, err := in.GetDesktopEnvironment(); err == nil {
			fmt.Printf("%s %s %s\n", Green.Sprint("desktop env"), "~", deskEnv)
		}
		waitGroup.Done()
	}()
	go func() {
		if terminal, err := in.GetTerminalInfo(); err == nil {
			fmt.Printf("%s %s %s\n", Cyan.Sprint("terminal"), "~", terminal)
		}
		waitGroup.Done()
	}()
	go func() {
		if cpu, err := in.GetCPU(); err == nil {
			fmt.Printf("%s %s %s\n", Blue.Sprint("cpu"), "~", cpu)
		}
		waitGroup.Done()
	}()
	go func() {
		if gpu, err := in.GetGPU(); err == nil {
			fmt.Printf("%s %s %s\n", Yellow.Sprint("gpu"), "~", gpu)
		}
		waitGroup.Done()
	}()
	go func() {
		if memoryUsage, err := in.GetMemoryUsage(); err == nil {
			fmt.Printf("%s %s %s\n", Red.Sprint("memory"), "~", memoryUsage)
		}
		waitGroup.Done()
	}()

	waitGroup.Wait()
	// Dots
	fmt.Printf("\n%s", Red.Sprint("○"))
	fmt.Printf("     %s", Green.Sprint("○"))
	fmt.Printf("     %s", Blue.Sprint("○"))
	fmt.Printf("     %s", Yellow.Sprint("○"))
	fmt.Printf("     %s", Cyan.Sprint("○"))
	fmt.Printf("     %s", Magenta.Sprint("○"))
	fmt.Printf("     %s\n", White.Sprint("○"))
}
