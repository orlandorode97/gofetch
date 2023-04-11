package windows

import (
	"os"
	"os/user"
)

var (
	getCurrent  = user.Current
	getHostname = os.Hostname
)

func (w *windows) GetName() string {
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
