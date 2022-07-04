package macos

import (
	"regexp"
	"strings"
)

var regexGPU *regexp.Regexp

// GetGPU returns the name of the GPU.
func (m *macos) GetGPU() string {
	regexGPU = regexp.MustCompile(`(:\s)(.*)`)

	cmd := "system_profiler SPDisplaysDataType | grep 'Chipset Model'"
	output, err := execCommand("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		return "Unknown"
	}
	gpu := strings.TrimSuffix(string(output), "\n")
	if regexGPU.MatchString(gpu) {
		gpu = regexGPU.FindStringSubmatch(gpu)[2]
	}
	gpu = strings.TrimLeft(gpu, " ")
	return gpu
}
