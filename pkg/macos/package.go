package macos

import (
	"fmt"
	"strings"
)

// GetNumberPackages return the number of packages installed by the current package manager.
func (m *macos) GetNumberPackages() string {
	cmd := "brew list | wc -l"
	output, err := execCommand("bash", "-c", cmd).Output()
	if err != nil {
		return "Unknown"
	}
	number := strings.TrimSuffix(string(output), "\n")

	number = strings.TrimSpace(number)
	return fmt.Sprintf("%s (%s)", number, "brew")
}
