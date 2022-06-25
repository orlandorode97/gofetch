package macos

import (
	"regexp"
	"strings"
)

var regexGPU *regexp.Regexp

// GetGPU returns the name of the GPU
func (m *macos) GetGPU() string {
	regexGPU = regexp.MustCompile(`(Intel|Advanced|NVIDIA|MCST|Virtual Box)([^\(|\(|\\]+)`)

	cmd := "system_profiler SPDisplaysDataType | grep 'Chipset Model'"
	output, err := execCommand("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		return "Unknown"
	}
	GPUs := strings.Split(string(output), "Chipset Model: ")
	GPU := strings.TrimSpace(GPUs[1])
	return GPU
}
