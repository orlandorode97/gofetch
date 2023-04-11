package windows

import (
	"fmt"
	"regexp"
)

var regexResolution *regexp.Regexp

func (w *windows) GetResolution() string {
	regexResolution = regexp.MustCompile(`\d+`)
	horizontal, err := getScreenValue("CurrentHorizontalResolution")
	if err != nil {
		return "Unknown"
	}

	vertical, err := getScreenValue("CurrentVerticalResolution")
	if err != nil {
		return "Unknown"
	}

	return fmt.Sprintf("%v X %v", *horizontal, *vertical)
}

func getScreenValue(arg string) (*string, error) {
	cmd := fmt.Sprintf("wmic path Win32_VideoController get %s", arg)
	output, err := execCommand("cmd", "/c", cmd).CombinedOutput()
	if err != nil {
		return nil, err
	}
	val := ""
	if regexResolution.MatchString(string(output)) {
		val = regexResolution.FindString(string(output))
	}
	return &val, nil
}
