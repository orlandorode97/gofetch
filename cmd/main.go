package cmd

import (
	"fmt"

	// macos "github.com/OrlandoRomo/gofetch/macos"

	"github.com/OrlandoRomo/gofetch/macos"
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

func InitCMD() {
	currentOS := macos.GetInfo()
	fmt.Printf("\t%s@%s\n", Red.Sprint(currentOS.Name), Red.Sprint(currentOS.Host))

	fmt.Printf("\t\t%s %s %s %s %s\n\n", Red.Sprint("X"), Green.Sprint("────"), Yellow.Sprint("X"), Green.Sprint("────"), Blue.Sprint("X"))
	fmt.Printf("\t%s %s %s\n", Red.Sprint("os"), "~", currentOS.OS)
	fmt.Printf("\t%s %s %s\n", Green.Sprint("host"), "~", currentOS.Host)
	fmt.Printf("\t%s %s %s\n", Cyan.Sprint("uptime"), "~", currentOS.Uptime)
	fmt.Printf("\t%s %s %s\n", Blue.Sprint("packages"), "~", currentOS.Packages)
	fmt.Printf("\t%s %s %s\n", Yellow.Sprint("shell"), "~", currentOS.Shell)
	fmt.Printf("\t%s %s %s\n", Red.Sprint("resolution"), "~", currentOS.Resolution)
	fmt.Printf("\t%s %s %s\n", Green.Sprint("desktop env"), "~", currentOS.DesktopEnvironment)
	fmt.Printf("\t%s %s %s\n", Cyan.Sprint("terminal"), "~", currentOS.Terminal)
	fmt.Printf("\t%s %s %s\n", Blue.Sprint("cpu"), "~", currentOS.CPU)
	fmt.Printf("\t%s %s %s\n", Yellow.Sprint("gpu"), "~", currentOS.GPU)
	fmt.Printf("\t%s %s %s\n\n", Red.Sprint("memory"), "~", currentOS.Memory)
	// Dots
	fmt.Printf("\t  %s", Red.Sprint("○"))
	fmt.Printf("     %s", Green.Sprint("○"))
	fmt.Printf("     %s", Blue.Sprint("○"))
	fmt.Printf("     %s", Yellow.Sprint("○"))
	fmt.Printf("     %s", Cyan.Sprint("○"))
	fmt.Printf("     %s", Magenta.Sprint("○"))
	fmt.Printf("     %s\n", White.Sprint("○"))
}
