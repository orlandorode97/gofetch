package linux

import (
	"os/exec"

	"github.com/orlandorode97/gofetch/fetch"
)

var execCommand = exec.Command

type Command string

type linux struct{}

func New() fetch.Fetcher {
	return &linux{}
}
