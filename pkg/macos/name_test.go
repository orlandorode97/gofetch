package macos

import (
	"errors"
	"os/user"
	"testing"
)

func TestGetName(t *testing.T) {
	tcs := []struct {
		Desc            string
		UserName        string
		Hostname        string
		MockCurrentUser func() (*user.User, error)
		MockHostname    func() (string, error)
		Expected        string
	}{
		{
			Desc:     "success - received macos name",
			UserName: "macos.user",
			Hostname: "macos.hostname",
			MockCurrentUser: func() (*user.User, error) {
				return &user.User{
					Username: "macos.user",
				}, nil
			},
			MockHostname: func() (string, error) {
				return "macos.hostname", nil
			},
			Expected: "macos.user@macos.hostname",
		},
		{
			Desc: "unable to get current user",
			MockCurrentUser: func() (*user.User, error) {
				return &user.User{}, errors.New("unable to get current user")
			},
			MockHostname: func() (string, error) {
				return "macos.hostname", nil
			},
			Expected: "Unknown@macos.hostname",
		},
		{
			Desc:     "unable to get hostname",
			UserName: "macos.user",
			MockCurrentUser: func() (*user.User, error) {
				return &user.User{
					Username: "macos.user",
				}, nil
			},
			MockHostname: func() (string, error) {
				return "", errors.New("unable to get hostname")
			},
			Expected: "macos.user@Unknown",
		},
	}

	for _, tt := range tcs {
		getCurrent, getHostname = tt.MockCurrentUser, tt.MockHostname
		t.Run(tt.Desc, func(t *testing.T) {
			mac := New()
			name := mac.GetName()
			if name != tt.Expected {
				t.Fatalf("received %s but expected %s", name, tt.Expected)
			}
		})
	}
}
