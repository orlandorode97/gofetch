package main

import (
	"fmt"

	"github.com/fatih/color"
)

type OSInformer interface {
}

type OSInformation struct {
	Name               string
	Host               string
	Uptime             string
	Packages           string
	Shell              string
	Resolution         string
	DesktopEnvironment string
	Terminal           string
	CPU                string
	GPU                string
	Memory             string
}

func main() {
	currentOS := GetInfo()

	fmt.Printf("\t\t\t\t %s %s\n", color.RedString("Name:"), currentOS.Name)
	fmt.Printf("\t\t\t\t %s %s\n", color.RedString("Host:"), currentOS.Host)
	fmt.Printf("\t\t\t\t %s %s\n", color.RedString("Uptime:"), currentOS.Uptime)
	fmt.Printf("\t\t\t\t %s %s\n", color.RedString("Packages:"), currentOS.Packages)
	fmt.Printf("\t\t\t\t %s %s\n", color.RedString("Shell:"), currentOS.Shell)
	fmt.Printf("\t\t\t\t %s %s\n", color.RedString("Resolution:"), currentOS.Resolution)
	fmt.Printf("\t\t\t\t %s %s\n", color.RedString("Terminal:"), currentOS.Terminal)
	fmt.Printf("\t\t\t\t %s %s\n", color.RedString("CPU:"), currentOS.CPU)
	fmt.Printf("\t\t\t\t %s %s\n", color.RedString("GPU:"), currentOS.GPU)
	fmt.Printf("\t\t\t\t %s %s\n", color.RedString("Memory:"), currentOS.Memory)
}
