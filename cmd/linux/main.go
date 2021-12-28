package main

import (
	"github.com/OrlandoRomo/gofetch"
	"github.com/OrlandoRomo/gofetch/command"
)

func main() {
	linux := gofetch.NewLinux()
	command.Fetch(linux)
}
