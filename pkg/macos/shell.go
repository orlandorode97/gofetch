package macos

import (
	"os"
	"regexp"
	"strings"
)

var (
	regexpShellVersion *regexp.Regexp
	regexpShell        *regexp.Regexp
)

// GetShellInformation return the used shell and its version.
func (m *macos) GetShellInformation() string {
	// regex to match /bin/bash /bin/zsh etc.
	regexpShell = regexp.MustCompile(`.*/(bash|zsh|fish|fish|csh|tcsh|ksh)`)
	// regex to match shell version like 1.5.6, 90.4, 1.2, 5.
	regexpShellVersion = regexp.MustCompile(`(\d.*\d).*`)

	output, err := execCommand(os.ExpandEnv("$SHELL"), "--version").CombinedOutput()
	if err != nil {
		return "Unknown"
	}
	return strings.TrimSuffix(string(output), "\n")
}
