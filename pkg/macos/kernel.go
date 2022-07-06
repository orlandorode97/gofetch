package macos

import (
	"regexp"
	"strings"
)

var regexKernel *regexp.Regexp

func (m *macos) GetKernelVersion() string {
	regexKernel = regexp.MustCompile(`^[^:]*`)
	output, err := execCommand("uname", "-v").CombinedOutput()
	if err != nil {
		return "Unknown"
	}

	kernel := strings.TrimSuffix(string(output), "\n")
	if !regexKernel.MatchString(kernel) {
		return "Unknown"
	}

	kernel = regexKernel.FindString(kernel)

	return kernel
}
