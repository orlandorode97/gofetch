package macos

import (
	"os"
	"strings"
)

// GetTerminalInfo get the current terminal name.
func (m *macos) GetTerminalInfo() string {
	output, err := execCommand("echo", os.ExpandEnv("$TERM_PROGRAM")).CombinedOutput()
	if err != nil {
		return "Unknown"
	}
	return strings.TrimSpace(string(output))
}
