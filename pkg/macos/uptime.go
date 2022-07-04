package macos

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/orlandorode97/gofetch/pkg/time"
)

var regexUptime *regexp.Regexp

// GetUptime returns the up time of the current OS.
func (m *macos) GetUptime() string {
	regexUptime = regexp.MustCompile(`(\d.*),`)

	output, err := execCommand("sysctl", "-n", "kern.boottime").CombinedOutput()
	if err != nil {
		return "Unknown"
	}

	bootTime := strings.TrimSuffix(string(output), "\n")
	if !regexUptime.MatchString(bootTime) {
		return "Unknown"
	}

	bootTime = regexUptime.FindStringSubmatch(bootTime)[1]
	now := `$(date +%s)`
	seconds := fmt.Sprintf("echo $((%s - %s))", now, bootTime)
	output, err = execCommand("bash", "-c", seconds).CombinedOutput()
	if err != nil {
		return "Unknown"
	}

	uptime := strings.TrimSuffix(string(output), "\n")
	return time.ParseUptime(uptime)
}
