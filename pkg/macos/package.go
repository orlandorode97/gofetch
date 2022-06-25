package macos

import (
	"strings"
)

// GetNumberPackages return the number of packages installed by the current package manager.
func (m *macos) GetNumberPackages() string {
	cmd := "ls /usr/local/Cellar/ | wc -l"
	output, err := execCommand("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		return "Unknown"
	}
	return strings.TrimSpace(string(output))
}
