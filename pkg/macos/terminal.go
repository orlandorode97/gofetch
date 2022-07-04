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
	shell := strings.TrimSuffix(string(output), "\n")
	if shell != "" {
		return shell
	}
	// Some shells dont't bind the env variable $TERM_PROGRAM, then is possible to use only $TERM env variable
	output, err = execCommand("echo", os.ExpandEnv("$TERM")).CombinedOutput()
	if err != nil {
		return "Unknown"
	}
	return strings.TrimSuffix(string(output), "\n")
}
