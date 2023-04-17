package windows

import (
	"regexp"
	"strings"
)

var regexOSVersion *regexp.Regexp

func (w *windows) GetOSVersion() string {
	regexOSVersion = regexp.MustCompile(`\n(.*)`)
	cmd := "wmic os get caption"
	output, err := execCommand("cmd", "/c", cmd).CombinedOutput()
	if err != nil {
		return "Unknown"
	}

	osVersion := ""
	if regexOSVersion.MatchString(string(output)) {
		osVersion = regexOSVersion.FindString(string(output))
		osVersion = strings.TrimLeft(osVersion, "\n")
	}

	return osVersion
}
