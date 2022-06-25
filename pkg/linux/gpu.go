package linux

import (
	"regexp"
	"strings"
)

var regexGPU *regexp.Regexp

// GetGPU returns the name of the GPU.
func (l *linux) GetGPU() string {
	regexGPU = regexp.MustCompile(`(:\s)(.*?)(\(|\s$|$)`)

	cmd := "lspci -v | grep 'VGA\\|Display\\|3D'"
	output, err := execCommand("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		return "Unknown"
	}

	gpu := strings.TrimSuffix(string(output), "\n")
	if !regexGPU.MatchString(gpu) {
		return "Unknown"
	}

	if regexGPU.MatchString(gpu) {
		gpu = regexGPU.FindStringSubmatch(gpu)[2]
		// Remove extra spaces at the end
		gpu = strings.TrimRight(gpu, " ")
	}

	return gpu
}
