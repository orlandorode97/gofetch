package main

import (
	"github.com/OrlandoRomo/gofetch/command"
	"github.com/OrlandoRomo/gofetch/linux"
)

func main() {
	linux := linux.New()
	command.PrintInfo(linux)
}
