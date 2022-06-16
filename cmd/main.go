package main

import (
	"runtime"

	"github.com/OrlandoRomo/gofetch/command"
	"github.com/OrlandoRomo/gofetch/linux"
	"github.com/OrlandoRomo/gofetch/macos"
	"github.com/OrlandoRomo/gofetch/windows"
)


func main(){
	var os command.Informer
	switch  goos:= runtime.GOOS; goos {
		case "darwin":
			os = macos.New()
		case "linux":
			os = linux.New()
		default:
			os = windows.New()
	}
	command.Fetch(os)
}
