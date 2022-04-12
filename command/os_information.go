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
	Red         *color.Color
	Green       *color.Color
	Cyan        *color.Color
	Yellow      *color.Color
	Blue        *color.Color
	Magenta     *color.Color
	White       *color.Color
	Colors      []*color.Color
	ShortenInfo map[string]string
)

func init() {
	Red = color.New(color.FgRed, color.Bold)
	Green = color.New(color.FgGreen, color.Bold)
	Cyan = color.New(color.FgCyan, color.Bold)
	Yellow = color.New(color.FgYellow, color.Bold)
	Blue = color.New(color.FgBlue, color.Bold)
	Magenta = color.New(color.FgHiMagenta, color.Bold)
	White = color.New(color.FgWhite, color.Bold)

	Colors = []*color.Color{Red, Green, Cyan, Yellow, Blue, Magenta}
	ShortenInfo = map[string]string{
		"GetOSVersion":          "OS",
		"GetName":               "name",
		"GetHostname":           "host",
		"GetUptime":             "uptime",
		"GetNumberPackages":     "packages",
		"GetShellInformation":   "shell",
		"GetResolution":         "resolution",
		"GetDesktopEnvironment": "DE",
		"GetTerminalInfo":       "terminal",
		"GetGPU":                "GPU",
		"GetCPU":                "CPU",
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

	waitGroup := sync.WaitGroup{}
	var m sync.Mutex
	inType, inValue := reflect.TypeOf(in), reflect.ValueOf(in)
	str := make([]string, 0)
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
			m.Lock()
			defer m.Unlock()
			str = append(str, fmt.Sprintf("%s %s %s", randColor(function.Name), "~", output))
		}(function, &waitGroup)
	}

	waitGroup.Wait()
	printWithLogo(str)
}

func randColor(s string) string {
	l := len(Colors)
	index := rand.Intn(l-0) + 0
	return Colors[index].Sprint(ShortenInfo[s])
}

func printWithLogo(infos []string) {
	beginning := fmt.Sprintf("%s %s %s %s %s\n\n", Red.Sprint("X"), Green.Sprint("────"), Yellow.Sprint("X"), Green.Sprint("────"), Blue.Sprint("X"))
	ending := fmt.Sprintf("%s %s %s %s %s %s %s\n",
		Red.Sprint("○"),
		Green.Sprint("○"),
		Blue.Sprint("○"),
		Yellow.Sprint("○"),
		Cyan.Sprint("○"),
		Magenta.Sprint("○"),
		White.Sprint("○"),
	)
	logo := fmt.Sprintf("%s %s\n %s\n %s\n %s\n %s\n %s\n %s\n %s\n %s\n %s\n %s\n %s\n %s",
		beginning,
		infos[0],
		infos[1],
		infos[2],
		infos[3],
		infos[4],
		infos[5],
		infos[6],
		infos[7],
		infos[8],
		infos[9],
		infos[10],
		infos[11],
		ending,
	)
	fmt.Print(logo)
}
