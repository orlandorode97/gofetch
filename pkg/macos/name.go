package macos

import (
	"os"
	"os/user"
)

// GetName returns the current user name.
func (m *macos) GetName() string {
	user, err := user.Current()
	if err != nil {
		user.Username = "Unknown"
	}

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "Unknown"
	}

	return user.Username + "@" + hostname
}
