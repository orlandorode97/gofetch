package main

import (
	"fmt"
	"time"
)

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
	now := time.Now()
	currentOS := GetInfo()
	fmt.Printf("Name: %s\n", currentOS.Name)
	fmt.Printf("Host: %s\n", currentOS.Host)
	fmt.Printf("Uptime: %s\n", currentOS.Uptime)
	fmt.Printf("Packages: %s\n", currentOS.Packages)
	fmt.Printf("Shell: %s\n", currentOS.Shell)
	duration := time.Since(now)
	fmt.Println("Took", duration, "to finished")

}
