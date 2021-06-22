package main

import (
	"fmt"

	macos "github.com/OrlandoRomo/go-fetch/macos"
	"github.com/fatih/color"
)

func main() {
	currentOS := macos.GetInfo()
	fmt.Printf("Welcome ~ %s@%s\n\n", color.RedString(currentOS.Name), color.RedString(currentOS.Host))
	fmt.Printf("\t%s %s %s\n", color.RedString("os"), "~", currentOS.OS)
	fmt.Printf("\t%s %s %s\n", color.RedString("host"), "~", currentOS.Host)
	fmt.Printf("\t%s %s %s\n", color.RedString("uptime"), "~", currentOS.Uptime)
	fmt.Printf("\t%s %s %s\n", color.RedString("packages"), "~", currentOS.Packages)
	fmt.Printf("\t%s %s %s\n", color.RedString("shell"), "~", currentOS.Shell)
	fmt.Printf("\t%s %s %s\n", color.RedString("resolution"), "~", currentOS.Resolution)
	fmt.Printf("\t%s %s %s\n", color.RedString("desktop env"), "~", currentOS.DesktopEnvironment)
	fmt.Printf("\t%s %s %s\n", color.RedString("terminal"), "~", currentOS.Terminal)
	fmt.Printf("\t%s %s %s\n", color.RedString("cpu"), "~", currentOS.CPU)
	fmt.Printf("\t%s %s %s\n", color.RedString("gpu"), "~", currentOS.GPU)
	fmt.Printf("\t%s %s %s\n", color.RedString("memory"), "~", currentOS.Memory)
}
