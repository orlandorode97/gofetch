package linux

import (
	"fmt"
	"strings"

	"github.com/orlandorode97/gofetch/pkg/time"
)

// GetUptime returns the up time of the current OS.
func (l *linux) GetUptime() string {
	boot := `$(date -d "$(uptime -s)" +%s)`
	now := `$(date +%s)`
	seconds := fmt.Sprintf("echo $((%s - %s))", now, boot)
	output, err := execCommand("bash", "-c", seconds).CombinedOutput()
	if err != nil {
		return "Unknown"
	}

	uptime := strings.TrimSuffix(string(output), "\n")
	return time.ParseUptime(uptime)
}
