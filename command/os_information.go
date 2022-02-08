package command

import (
	"fmt"
	"math/rand"
	"os/exec"
	"reflect"
	"strings"
	"sync"

	"github.com/fatih/color"
)

var (
	Red     *color.Color
	Green   *color.Color
	Cyan    *color.Color
	Yellow  *color.Color
	Blue    *color.Color
	Magenta *color.Color
	White   *color.Color
	Colors  []*color.Color
	Info    map[string]string
)

func init() {
	Red = color.New(color.FgRed, color.Bold)
	Green = color.New(color.FgGreen, color.Bold)
	Cyan = color.New(color.FgCyan, color.Bold)
	Yellow = color.New(color.FgYellow, color.Bold)
	Blue = color.New(color.FgBlue, color.Bold)
	Magenta = color.New(color.FgHiMagenta, color.Bold)
	White = color.New(color.FgWhite, color.Bold)

	Colors = []*color.Color{Red, Green, Cyan, Yellow, Blue, Magenta, White}
	Info = map[string]string{
		"GetOSVersion":          "os",
		"GetName":               "name",
		"GetHostname":           "host",
		"GetUptime":             "uptime",
		"GetNumberPackages":     "packages",
		"GetShellInformation":   "shell",
		"GetResolution":         "resolution",
		"GetDesktopEnvironment": "de",
		"GetTerminalInfo":       "terminal",
		"GetGPU":                "gpu",
		"GetCPU":                "cpu",
		"GetMemoryUsage":        "memory",
	}
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
	bufErr.Reset()
	bufOut.Reset()

	cmd := exec.Command(command, args...)

	cmd.Stderr = &bufErr.stderr
	cmd.Stdout = &bufOut.stdout

	err := cmd.Run()

	return strings.TrimSuffix(bufOut.stdout.String(), "\n"), err
}

func Fetch(in Informer) {
	fmt.Printf("%s %s %s %s %s\n\n", Red.Sprint("X"), Green.Sprint("────"), Yellow.Sprint("X"), Green.Sprint("────"), Blue.Sprint("X"))

	waitGroup := sync.WaitGroup{}
	inType := reflect.TypeOf(in)
	inValue := reflect.ValueOf(in)
	for i := 0; i < inType.NumMethod(); i++ {
		function := inType.Method(i)
		waitGroup.Add(1)
		go func(f reflect.Method, w *sync.WaitGroup) {
			defer waitGroup.Done()
			result := function.Func.Call([]reflect.Value{inValue})
			output, _ := result[0].Interface().(string)
			err, ok := result[1].Interface().(error)
			if ok && err != nil {
				return
			}
			fmt.Printf("%s %s %s\n", RandColor(function.Name), "~", output)
		}(function, &waitGroup)
	}
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

func RandColor(s string) string {
	l := len(Colors)
	index := rand.Intn(l-0) + 0
	return Colors[index].Sprint(Info[s])
}
