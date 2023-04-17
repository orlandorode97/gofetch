package windows

import (
	"fmt"
	"strings"
)

func (w *windows) GetNumberPackages() string {
	cmd := "where scoop"
	result, err := countPackages(cmd, "scoop")
	if err == nil {
		return result
	}

	cmd = "where choco"
	result, err = countPackages(cmd, "choco")
	if err == nil {
		return result
	}

	return "Unknown"
}

func countPackages(cmd, pkgManagerName string) (string, error) {
	_, err := execCommand("cmd", "/c", cmd).CombinedOutput()
	if err != nil {
		return "Unknown", err
	}

	totalCmd := fmt.Sprintf(`(ls C:\ProgramData\%s | measure-object -line).Lines`, pkgManagerName)
	if pkgManagerName == "scoop" {
		totalCmd = `(ls ~/scoop/apps/* | measure-object -line).Lines`
	}

	output, err := execCommand("powershell", "-nologo", "-noprofile", totalCmd).CombinedOutput()
	if err != nil {
		return "Unknown", err
	}

	total := strings.TrimSuffix(string(output), "\r\n")

	return fmt.Sprintf("%s (%s)", total, pkgManagerName), nil
}
