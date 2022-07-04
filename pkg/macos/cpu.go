package macos

import (
	"regexp"
	"strings"
)

var regexCPU *regexp.Regexp

// GetCPU returns the name of th CPU.
func (m *macos) GetCPU() string {
	regexCPU = regexp.MustCompile(`\s.*`)
	cmd := "sysctl -a | grep machdep.cpu.brand_string"
	output, err := execCommand("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		return "Unknown"
	}

	cpu := strings.TrimSuffix(string(output), "\n")
	if regexCPU.MatchString(cpu) {
		cpu = regexCPU.FindAllString(cpu, -1)[0]
	}

	cpu = strings.TrimLeft(cpu, " ")
	return cpu
}
