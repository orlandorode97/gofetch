package macos

import (
	"os/exec"

	"github.com/orlandorode97/gofetch/fetch"
)

var execCommand = exec.Command

type macos struct{}

func New() fetch.Fetcher {
	return &macos{}
}
