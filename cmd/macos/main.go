package main

import (
	"github.com/OrlandoRomo/gofetch/command"
	"github.com/OrlandoRomo/gofetch/macos"
)

func main() {
	mac := macos.New()
	command.PrintInfo(mac)
}
