package linux

import (
	"os"
	"os/user"
)

var (
	getCurrent  = user.Current
	getHostname = os.Hostname
)

// GetName returns the current user name.
func (l *linux) GetName() string {
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
