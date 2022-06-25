package linux

import (
	"regexp"
	"strings"
)

var regexpCPU *regexp.Regexp

// GetCPU returns the name of th CPU.
func (l *linux) GetCPU() string {
	// Model name:        Intel(R) Core(TM) i5-1020U => Intel(R) Core(TM) i5-1020U.
	// BIOS Model name:        Intel(R) Core(TM) i5-1020U => Will not match since we care only for the previous grep result.
	regexpCPU = regexp.MustCompile(`^(Model name:.*\s{2})([A-Z].*)`)

	cmd := "lscpu | grep 'Model name:'"
	output, err := execCommand("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		return "Unknown"
	}

	cpu := strings.TrimSuffix(string(output), "\n")
	if regexpCPU.MatchString(cpu) {
		cpu = regexpCPU.FindStringSubmatch(cpu)[2]
	}

	return cpu
}
