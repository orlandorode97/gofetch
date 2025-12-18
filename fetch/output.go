package fetch

import (
	"fmt"
	"math/rand"
	"reflect"
	"sync"

	"github.com/fatih/color"
)

var (
	Red     = color.New(color.FgRed, color.Bold)
	Green   = color.New(color.FgGreen, color.Bold)
	Cyan    = color.New(color.FgCyan, color.Bold)
	Yellow  = color.New(color.FgYellow, color.Bold)
	Blue    = color.New(color.FgBlue, color.Bold)
	Magenta = color.New(color.FgHiMagenta, color.Bold)
	White   = color.New(color.FgWhite, color.Bold)
	Colors  []*color.Color
	fields  map[string]string
)

func init() {
	initColorFields()
}

var gopher = `
%s
⠀⠀⠀⠀⠀⠀⠀⠀⢀⣤⣶⣾⣿⣿⣿⣿⣿⣿⣶⣦⣄⠀⠀⠀⠀⠀⠀ %s
⠀⠀⠀⠀⢠⡶⣦⣴⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣦⡴⣦⠀⠀ %s
⠀⠀⠀⠀⠀⠙⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟⠉⠀⠀ %s
⠀⠀⠀⠀⠀⠀⣿⣿⣿⣟⠁⠊⣿⣿⣿⣿⣿⡏⠒⠈⣿⣿⣿⡇⠀⠀⠀ %s
⠀⠀⠀⠀⠀⠀⢿⣿⣿⣿⣷⣾⣿⣿⠿⠿⣿⣿⣶⣾⣿⣿⣿⡇⠀⠀⠀ %s
⠀⠀⠀⠀⠀⠀⢸⣿⣿⣿⣿⣿⣿⣷⣤⣤⣿⣿⣿⣿⣿⣿⣿⣧⣼⠛⠀ %s
⠀⠀⠀⠀⠀⠀⣸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠃⠀⠀ %s
⠀⠀⠀⠀⠀⣰⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡏⠀⠀⠀ %s
⠀⠀⠀⠴⡾⠋⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀ %s
⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀ %s
⠀⠀⠀⠀⠀⠀⢹⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠇⠀⠀⠀ %s
⠀⠀⠀⠀⠀⠀⠀⢻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠏⠀⠀⠀⠀ %s
⠀⠀⠀⠀⠀⠀⠀⠀⣹⣿⠿⣿⣿⣿⣿⣿⣿⣿⣿⠿⣿⡁⠀⠀⠀⠀⠀
%s`

func initColorFields() {
	Colors = []*color.Color{Red, Green, Cyan, Yellow, Blue, Magenta}

	fields = map[string]string{
		"GetOSVersion":          "OS",
		"GetName":               "name",
		"GetKernelVersion":      "kernel",
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

func Fetch(in Fetcher) {
	w := sync.WaitGroup{}
	var m sync.Mutex
	inType, inValue := reflect.TypeOf(in), reflect.ValueOf(in)
	outputFields := make(map[string]string, 0)
	for i := 0; i < inType.NumMethod(); i++ {
		function := inType.Method(i)
		w.Add(1)
		go func(f reflect.Method) {
			defer w.Done()
			result := f.Func.Call([]reflect.Value{inValue})
			output, _ := result[0].Interface().(string)
			m.Lock()
			name := fields[f.Name]
			outputFields[name] = colorOutput(output, name)
			m.Unlock()
		}(function)
	}

	w.Wait()

	dots := fmt.Sprintf("%s  	%s  	%s  	%s  	%s  	%s  	%s  	\n",
		Red.Sprint("○"),
		Green.Sprint("○"),
		Blue.Sprint("○"),
		Yellow.Sprint("○"),
		Cyan.Sprint("○"),
		Magenta.Sprint("○"),
		White.Sprint("○"),
	)
	fmt.Printf(
		Cyan.Sprint(gopher),
		dots,
		outputFields["name"],
		outputFields["OS"],
		outputFields["kernel"],
		outputFields["uptime"],
		outputFields["packages"],
		outputFields["shell"],
		outputFields["resolution"],
		outputFields["DE"],
		outputFields["terminal"],
		outputFields["CPU"],
		outputFields["GPU"],
		outputFields["memory"],
		dots,
	)
}

func colorOutput(output, fieldName string) string {
	color := RandColor()
	return fmt.Sprintf("%s %s %s", color.Sprint(fieldName), "~", output)
}

func RandColor() *color.Color {
	l := len(Colors)
	index := rand.Intn(l)
	return Colors[index]
}
