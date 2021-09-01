package cmd

import (
	"runtime"

	"github.com/OrlandoRomo/gofetch/linux"
	"github.com/OrlandoRomo/gofetch/macos"
	"github.com/OrlandoRomo/gofetch/os"
	"github.com/OrlandoRomo/gofetch/windows"
)

func InitCMD() {
	switch runtime.GOOS {
	case "darwin":
		mac := macos.NewMacOS()
		os.PrintInfo(mac)
	case "linux":
		lx := linux.NewLinux()
		os.PrintInfo(lx)
	case "windows":
		ws := windows.NewWindows()
		os.PrintInfo(ws)

	}
}
