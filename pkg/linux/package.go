package linux

import (
	"fmt"
	"regexp"
	"strings"
)

// Command to found the binary file of the current package manager.
const NetPackage = `which {xbps-install,apk,dpkg,pacman,nix,yum,rpm,emerge} 2>/dev/null | grep -v "not found"`

type PackageManager string

var (
	regexPkgCmd  *regexp.Regexp
	regexNumbers *regexp.Regexp
	pkgManagers  map[PackageManager]Command
)

func initPkgCommands() {
	pkgManagers = map[PackageManager]Command{
		"xbps-install": "xbps-query -l | wc -l",
		"apk":          "apk search | wc -l",
		"dpkg":         "dpkg-query -f '.\n' -W | wc -l",
		"pacman":       "pacman -Q | wc -l",
		"nix":          `nix-env -qa --installed "*" | wc -l`,
		"yum":          "yum list installed | wc -l",
		"rpm":          "rpm -qa | wc -l",
		"emerge":       "qlist -I | wc -l",
	}
	// Regexg to match package manager name from inputs like /usr/bin/dpkg.
	regexPkgCmd = regexp.MustCompile(`[^/]*$`)
	// Some pkg managers command return warnings, errors, etc from stardard ouput, this regular expression is for capturing the number of packages.
	regexNumbers = regexp.MustCompile(`\d+`)
}

// GetNumberPackages return the number of packages installed by the current package manager.
func (l *linux) GetNumberPackages() string {
	initPkgCommands()

	output, err := execCommand("bash", "-c", NetPackage).CombinedOutput()
	if len(output) == 0 || err != nil {
		return "Unknown"
	}
	pkgManager := strings.TrimSuffix(string(output), "\n")

	if !regexPkgCmd.MatchString(pkgManager) {
		return "Unknown"
	}

	pkgManager = regexPkgCmd.FindString(pkgManager)

	name, ok := pkgManagers[PackageManager(pkgManager)]

	if !ok {
		return "Unknown"
	}

	output, err = execCommand("bash", "-c", string(name)).CombinedOutput()
	if err != nil {
		return "Unknown"
	}
	total := strings.TrimSuffix(string(output), "\n")

	if !regexNumbers.MatchString(total) {
		return "Unknown"
	}
	total = regexNumbers.FindString(total)

	return fmt.Sprintf("%s (%s)", total, pkgManager)
}
