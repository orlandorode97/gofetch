package os

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

type OSInformer interface {
	GetOSVersion() (string, error)
	GetName() (string, error)
	GetHostname() (string, error)
	GetUptime() (string, error)
	GetNumberPackages() (string, error)
	GetShellInformation() (string, error)
	GetResolution() (string, error)
	GetDesktopEnvironment() (string, error)
	GetTerminalInfo() (string, error)
	GetCPU() (string, error)
	GetGPU() (string, error)
	GetMemoryUsage() (string, error)
}

//ExecuteCommand executes the command with arguments as well reset the stderr and stdout
func ExecuteCommand(command string, args ...string) (string, error) {
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

func PrintInfo(informer OSInformer) {

	waitGroup := sync.WaitGroup{}

	waitGroup.Add(9)
	if name, err := informer.GetName(); err == nil {
		fmt.Printf("%s", Red.Sprint(name))
	}
	if host, err := informer.GetHostname(); err == nil {
		fmt.Printf("@%s\n", Red.Sprint(host))
	}
	fmt.Printf("%s %s %s %s %s\n\n", Red.Sprint("X"), Green.Sprint("────"), Yellow.Sprint("X"), Green.Sprint("────"), Blue.Sprint("X"))

	go func() {
		if uptime, err := informer.GetUptime(); err == nil {
			fmt.Printf("%s %s %s\n", Cyan.Sprint("uptime"), "~", uptime)
		}
		waitGroup.Done()
	}()
	go func() {
		if numPackages, err := informer.GetNumberPackages(); err == nil {
			fmt.Printf("%s %s %s\n", Blue.Sprint("packages"), "~", numPackages)
		}
		waitGroup.Done()
	}()
	go func() {
		if shell, err := informer.GetShellInformation(); err == nil {
			fmt.Printf("%s %s %s\n", Yellow.Sprint("shell"), "~", shell)
		}
		waitGroup.Done()
	}()
	go func() {
		if resolution, err := informer.GetResolution(); err == nil {
			fmt.Printf("%s %s %s\n", Red.Sprint("resolution"), "~", resolution)
		}
		waitGroup.Done()
	}()
	go func() {
		if deskEnv, err := informer.GetDesktopEnvironment(); err == nil {
			fmt.Printf("%s %s %s\n", Green.Sprint("desktop env"), "~", deskEnv)
		}
		waitGroup.Done()
	}()
	go func() {
		if terminal, err := informer.GetTerminalInfo(); err == nil {
			fmt.Printf("%s %s %s\n", Cyan.Sprint("terminal"), "~", terminal)
		}
		waitGroup.Done()
	}()
	go func() {
		if cpu, err := informer.GetCPU(); err == nil {
			fmt.Printf("%s %s %s\n", Blue.Sprint("cpu"), "~", cpu)
		}
		waitGroup.Done()
	}()
	go func() {
		if gpu, err := informer.GetGPU(); err == nil {
			fmt.Printf("%s %s %s\n", Yellow.Sprint("gpu"), "~", gpu)
		}
		waitGroup.Done()
	}()
	go func() {
		if memoryUsage, err := informer.GetMemoryUsage(); err == nil {
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
