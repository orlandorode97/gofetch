package macos

import (
	"os"
	"os/user"
)

var (
	getCurrent  = user.Current
	getHostname = os.Hostname
)

// GetName returns the current user name.
func (m *macos) GetName() string {
	user, err := getCurrent()
	if err != nil {
		user.Username = "Unknown"
	}

	hostname, err := getHostname()
	if err != nil {
		hostname = "Unknown"
	}

	return user.Username + "@" + hostname
}
