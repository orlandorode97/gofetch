package cmd

import (
	"fmt"

	// macos "github.com/OrlandoRomo/gofetch/macos"
	"github.com/OrlandoRomo/gofetch/linux"
	"github.com/fatih/color"
)

func InitCMD() {
	currentOS := linux.GetInfo()
	fmt.Printf("\nWelcome ~ %s@%s\n\n", color.RedString(currentOS.Name), color.RedString(currentOS.Host))
	fmt.Printf("\t%s %s %s\n", color.RedString("os"), "~", currentOS.OS)
	fmt.Printf("\t%s %s %s\n", color.GreenString("host"), "~", currentOS.Host)
	fmt.Printf("\t%s %s %s\n", color.CyanString("uptime"), "~", currentOS.Uptime)
	fmt.Printf("\t%s %s %s\n", color.BlueString("packages"), "~", currentOS.Packages)
	fmt.Printf("\t%s %s %s\n", color.YellowString("shell"), "~", currentOS.Shell)
	fmt.Printf("\t%s %s %s\n", color.RedString("resolution"), "~", currentOS.Resolution)
	fmt.Printf("\t%s %s %s\n", color.GreenString("desktop env"), "~", currentOS.DesktopEnvironment)
	fmt.Printf("\t%s %s %s\n", color.CyanString("terminal"), "~", currentOS.Terminal)
	fmt.Printf("\t%s %s %s\n", color.BlueString("cpu"), "~", currentOS.CPU)
	fmt.Printf("\t%s %s %s\n", color.YellowString("gpu"), "~", currentOS.GPU)
	fmt.Printf("\t%s %s %s\n\n", color.RedString("memory"), "~", currentOS.Memory)
}
