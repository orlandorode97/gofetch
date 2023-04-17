package windows

import (
	"regexp"
	"strings"
)

var regexGPU *regexp.Regexp

func (w *windows) GetGPU() string {
	regexGPU = regexp.MustCompile(`\n(.*)`)
	cmd := "wmic path win32_VideoController get caption"
	output, err := execCommand("cmd", "/c", cmd).CombinedOutput()
	if err != nil {
		return "Unknown"
	}

	gpu := ""
	if regexGPU.MatchString(string(output)) {
		gpu = regexGPU.FindString(string(output))
		gpu = strings.TrimLeft(gpu, "\n")
	}

	return gpu
}
