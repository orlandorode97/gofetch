package macos

import (
	"strings"
)

// GetResolution returns the resolution of thee current monitor.
func (m *macos) GetResolution() string {
	cmd := "system_profiler SPDisplaysDataType  | grep Resolution"
	output, err := execCommand("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		return "Unknown"
	}
	resolution := string(output)
	resolutions := strings.Split(resolution, "Resolution: ")
	resolution = strings.TrimSpace(resolutions[1])
	return resolution
}
