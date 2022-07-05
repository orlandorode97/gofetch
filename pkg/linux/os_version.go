package linux

import (
	"fmt"
	"regexp"
	"strings"
)

var regexOS *regexp.Regexp

// GetHostname returns the hostname of the linux distro.
func (l *linux) GetOSVersion() string {
	regexOS = regexp.MustCompile(`(NAME|VERSION)=(.+)`)
	// To get os name and version let's use the follow command: grep -E -i -w 'VERSION|NAME' /etc/os-release.
	output, err := execCommand("grep", "-E -i -w", "/etc/os-release").CombinedOutput()
	if err != nil {
		return "Unknown"
	}

	osVersion := strings.TrimSuffix(string(output), "\n")

	if !regexOS.MatchString(osVersion) {
		return "Unknown"
	}

	matches := regexOS.FindStringSubmatch(osVersion)
	fmt.Println(matches)
	return strings.TrimSuffix(string(output), "\n")
}
