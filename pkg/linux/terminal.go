package linux

import (
	"os"
	"strings"
)

// GetTerminalInfo get the current terminal name.
func (l *linux) GetTerminalInfo() string {
	output, err := execCommand("echo", os.ExpandEnv("$TERM")).CombinedOutput()
	if err != nil {
		return "Unknown"
	}

	return strings.TrimSpace(string(output))
}
