package linux

import (
	"regexp"
	"strings"
)

// Command to found the binary file of the current package manager.
const NetPackage = `which {xbps-install,apk,dpkg,pacman,nix,yum,rpm,emerge} 2>/dev/null | grep -v "not found"`

type PackageManager string

var (
	regexPackages   *regexp.Regexp
	distrosPackages map[PackageManager]Command
)

func initPkgCommands() {
	distrosPackages = map[PackageManager]Command{
		"xbps-install": "xbps-query -l | wc -l",
		"apk":          "apk search | wc -l",
		"dpkg":         "dpkg-query -f '.\n' -W | wc -l",
		"pacman":       "pacman -Q | wc -l",
		"nix":          `nix-env -qa --installed "*" | wc -l`,
		"yum":          "yum list installed | wc -l",
		"rpm":          "rpm -qa | wc -l",
		"emerge":       "qlist -I | wc -l",
	}
	regexPackages = regexp.MustCompile(`[^/]*$`)
}

// GetNumberPackages return the number of packages installed by the current package manager.
func (l *linux) GetNumberPackages() string {
	initPkgCommands()

	output, err := execCommand("bash", "-c", NetPackage).CombinedOutput()
	if err != nil {
		return "Unknown"
	}
	pkgManager := strings.TrimSuffix(string(output), "\n")

	if regexPackages.MatchString(pkgManager) {
		pkgManager = regexPackages.FindString(pkgManager)
	}

	name, ok := distrosPackages[PackageManager(pkgManager)]

	if !ok {
		return "Unknown"
	}

	output, err = execCommand("bash", "-c", string(name)).CombinedOutput()
	if err != nil {
		return "Unknown"
	}
	total := strings.TrimSuffix(string(output), "\n")

	return total
}
