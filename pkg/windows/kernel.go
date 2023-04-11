package windows

import (
	"regexp"
	"strings"
)

var regexKernel *regexp.Regexp

func (w *windows) GetKernelVersion() string {
	regexKernel = regexp.MustCompile(`\n(.*)`)
	cmd := "wmic os get version"
	output, err := execCommand("cmd", "/c", cmd).CombinedOutput()
	if err != nil {
		return "Unknown"
	}

	kernel := ""
	if regexKernel.MatchString(string(output)) {
		kernel = regexKernel.FindString(string(output))
		kernel = strings.TrimLeft(kernel, "\n")
	}

	return kernel
}
