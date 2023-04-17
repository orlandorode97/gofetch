package windows

import (
	"regexp"

	"github.com/orlandorode97/gofetch/pkg/time"
)

var regexUptime *regexp.Regexp

func (w *windows) GetUptime() string {
	regexUptime = regexp.MustCompile(`\d+`)
	cmd := `(get-date) - (gcim Win32_OperatingSystem).LastBootUpTime | findstr "TotalSeconds"`
	output, err := execCommand("powershell", "-nologo", "-noprofile", cmd).CombinedOutput()
	if err != nil {
		return "Unknown"
	}

	uptime := ""
	if regexUptime.MatchString(string(output)) {
		uptime = regexUptime.FindString(string(output))
	}

	return time.ParseUptime(uptime)
}
