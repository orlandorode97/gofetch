package main

import (
	"github.com/OrlandoRomo/gofetch"
	"github.com/OrlandoRomo/gofetch/command"
)

func main() {
	win := gofetch.NewWin()
	command.Fetch(win)
}
