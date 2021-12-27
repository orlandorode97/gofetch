package main

import (
	"github.com/OrlandoRomo/gofetch/command"
	"github.com/OrlandoRomo/gofetch/windows"
)

func main() {
	win := windows.New()
	command.PrintInfo(win)

}
