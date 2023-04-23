package windows

import "regexp"

var regexDeskEnvironment *regexp.Regexp

func (w *windows) GetDesktopEnvironment() string {
	regexDeskEnvironment = regexp.MustCompile(`10|8`)
	cmd := "wmic os get caption"
	output, err := execCommand("cmd", "/C", cmd).CombinedOutput()
	if err != nil {
		return "Unknown"
	}

	if regexDeskEnvironment.MatchString(string(output)) {
		winVersion := regexDeskEnvironment.FindString(string(output))
		if winVersion == "10" {
			return "Fluent"
		}
		return "Metro"
	}
	return "Aero"
}
