package windows

import (
	"regexp"
	"strings"
)

var regexCPU *regexp.Regexp

func (w *windows) GetCPU() string {
	regexCPU = regexp.MustCompile(`\n(.*)`)
	cmd := "wmic cpu get name"
	output, err := execCommand("cmd", "/c", cmd).CombinedOutput()
	if err != nil {
		return "Unknown"
	}

	cpu := ""
	if regexCPU.MatchString(string(output)) {
		cpu = regexCPU.FindString(string(output))
		cpu = strings.TrimLeft(cpu, "\n")
	}

	return cpu
}
