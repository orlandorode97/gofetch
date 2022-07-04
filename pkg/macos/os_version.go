package macos

import (
	"strings"
)

// GetHostname returns the hostname of the linux distro.
func (m *macos) GetOSVersion() string {
	output, err := execCommand("sw_vers", "-productName").CombinedOutput()
	if err != nil {
		return "Unknown"
	}
	name := strings.TrimSuffix(string(output), "\n")

	output, err = execCommand("sw_vers", "-productVersion").CombinedOutput()
	if err != nil {
		return "Unknown"
	}
	version := strings.TrimSuffix(string(output), "\n")

	return name + " " + version
}
