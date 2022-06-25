package linux

import (
	"regexp"
)

var regexpResolution *regexp.Regexp

// GetResolution returns the resolution of thee current monitor.
func (l *linux) GetResolution() string {
	regexpResolution = regexp.MustCompile(`\d+x\d+.*`)

	cmd := "xdpyinfo | grep 'dimensions:'"
	output, err := execCommand("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		return "Unknown"
	}
	resolution := string(output)
	if regexpResolution.MatchString(string(output)) {
		resolution = regexpResolution.FindString(resolution)
	}

	return resolution
}
