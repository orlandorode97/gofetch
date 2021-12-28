package main

import (
	"github.com/OrlandoRomo/gofetch"
	"github.com/OrlandoRomo/gofetch/command"
)

func main() {
	macos := gofetch.NewMacOS()
	command.Fetch(macos)
}
