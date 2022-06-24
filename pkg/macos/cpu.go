package macos

import (
	"strings"
)

// GetCPU returns the name of th CPU.
func (m *macos) GetCPU() string {
	cmd := "sysctl -a | grep machdep.cpu.brand_string"
	output, err := execCommand("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		return "Unknown"
	}

	splitCPU := strings.Split(string(output), ": ")

	CPU := strings.Replace(splitCPU[1], "\n\r", "", -1)
	CPU = strings.TrimSpace(CPU)
	return CPU
}
