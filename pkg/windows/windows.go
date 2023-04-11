package windows

import (
	"os/exec"

	"github.com/orlandorode97/gofetch/fetch"
)

var execCommand = exec.Command

type Command string

type windows struct{}

func New() fetch.Fetcher {
	return &windows{}
}
