package macos

import (
	"regexp"
	"strings"
)

var regexResolution *regexp.Regexp

// GetResolution returns the resolution of thee current monitor.
func (m *macos) GetResolution() string {
	regexResolution = regexp.MustCompile(`\d+ x \d+.*`)
	cmd := "system_profiler SPDisplaysDataType  | grep Resolution"
	output, err := execCommand("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		return "Unknown"
	}
	resolution := strings.TrimSuffix(string(output), "\n")

	if regexResolution.MatchString(resolution) {
		resolution = regexResolution.FindStringSubmatch(resolution)[0]
	}

	return resolution
}
