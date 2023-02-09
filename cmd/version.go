package main

import (
	"fmt"
	"runtime"

	"github.com/orlandorode97/gofetch/fetch"
	"github.com/urfave/cli/v2"
)

var (
	buildTime       string
	lastCommit      string
	semanticVersion string
)

func newVersionCommand() *cli.Command {
	return &cli.Command{
		Name:   "version",
		Usage:  "gets current version of gofetch",
		Action: versionAction,
	}
}

func versionAction(c *cli.Context) error {
	color := fetch.RandColor()

	fmt.Printf("%s: %s\n", color.Sprint("Semmantic version"), semanticVersion)
	fmt.Printf("%s: %s\n", color.Sprint("Commit"), lastCommit)
	fmt.Printf("%s: %s\n", color.Sprint("Build date"), buildTime)
	fmt.Printf("%s: %s/%s\n", color.Sprint("System version"), runtime.GOARCH, runtime.GOOS)
	fmt.Printf("%s: %s\n", color.Sprint("Golang version"), runtime.Version())
	return nil
}
