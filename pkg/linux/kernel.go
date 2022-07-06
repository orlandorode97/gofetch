package linux

import "strings"

func (l *linux) GetKernelVersion() string {
	output, err := execCommand("uname", "-smr").CombinedOutput()
	if err != nil {
		return "Unknown"
	}

	return strings.TrimSuffix(string(output), "\n")
}
