package main

import (
	"log"
	"os"
	"runtime"

	"github.com/orlandorode97/gofetch/fetch"
	"github.com/orlandorode97/gofetch/pkg/linux"
	"github.com/orlandorode97/gofetch/pkg/macos"
	"github.com/orlandorode97/gofetch/pkg/windows"
	"github.com/urfave/cli/v2"
)

func main() {
	gofetch := &cli.App{
		Commands: []*cli.Command{
			newVersionCommand(),
		},
		Name:   "gofetch",
		Usage:  "fetches os information",
		Action: Fetch,
	}
	if err := gofetch.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func Fetch(c *cli.Context) error {
	var os fetch.Fetcher
	switch goos := runtime.GOOS; goos {
	case "darwin":
		os = macos.New()
	case "linux":
		os = linux.New()
	default:
		os = windows.New()
	}

	fetch.Fetch(os)

	return nil
}
