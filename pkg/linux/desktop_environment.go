package linux

import (
	"os"
	"strings"
)

type DesktopName string

var desktopVersion map[DesktopName]Command

func initDesktops() {
	desktopVersion = map[DesktopName]Command{
		"Plasma":   "plasmashell --version",
		"KDE":      "plasmashell --version",
		"MATE":     "mate-session --version",
		"Xfce":     "xfce4-session --version",
		"GNOME":    "gnome-shell --version",
		"Cinnamon": "cinnamon --version",
		"Deepin":   "awk -F'=' '/MajorVersion/ {print $2}' /etc/os-version",
		"Budgie":   "budgie-desktop --version",
		"LXQt":     "lxqt-session --version",
		"Lumina":   "lumina-desktop --version 2>&1",
		"Trinity":  "tde-config --version",
		"Unity":    "unity --version",
	}
}

// GetDesktopEnvironment returns the resolution of the current monitor.
func (l *linux) GetDesktopEnvironment() string {
	initDesktops()

	// xdg stands for Cross-Desktop Group.
	output, err := execCommand("echo", os.ExpandEnv("$XDG_CURRENT_DESKTOP")).CombinedOutput()
	if err != nil {
		return "Unknown"
	}
	xdg := strings.TrimSuffix(string(output), "\n")
	// Some $XDG_CURRENT_DESKTOP values are like ubuntu:GNOME or simple values like Pantheon.
	deskName := strings.Split(xdg, ":")
	if len(deskName) != 2 {
		return deskName[0]
	}

	deVersionCommand := desktopVersion[DesktopName(deskName[1])]

	version, err := execCommand("bash", "-c", string(deVersionCommand)).CombinedOutput()
	if err != nil {
		return "Unknown"
	}

	return strings.TrimSuffix(string(version), "\n")
}
