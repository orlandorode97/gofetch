package linux

import "strings"

// GetHostname returns the hostname of the linux distro.
func (l *linux) GetOSVersion() string {
	output, err := execCommand("uname", "-srm").CombinedOutput()
	if err != nil {
		return "Unknown"
	}
	return strings.TrimSuffix(string(output), "\n")
}
