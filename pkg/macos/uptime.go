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
	regexUptime = regexp.MustCompile(`\=\s(.*.+?),`)

	boot, err := execCommand("sysctl", "-n", "kern.boottime").CombinedOutput()
	if err != nil {
		return "Unknown"
	}
	matches := regexUptime.FindStringSubmatch(string(boot))
	if len(matches) != 0 && matches[1] == "" {
		return "Unknown"
	}
	now := `$(date +%s)`
	seconds := fmt.Sprintf("echo $((%s - %s))", now, matches[1])
	output, err := execCommand("bash", "-c", seconds).CombinedOutput()
	if err != nil {
		return "Unknown"
	}

	uptime := strings.TrimSuffix(string(output), "\n")
	return time.ParseUptime(uptime)
}
