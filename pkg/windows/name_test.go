package windows

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
			Desc:     "success - received windows name",
			UserName: "windows.user",
			Hostname: "windows.hostname",
			MockCurrentUser: func() (*user.User, error) {
				return &user.User{
					Username: "windows.user",
				}, nil
			},
			MockHostname: func() (string, error) {
				return "windows.hostname", nil
			},
			Expected: "windows.user@windows.hostname",
		},
		{
			Desc: "unable to get current user",
			MockCurrentUser: func() (*user.User, error) {
				return &user.User{}, errors.New("unable to get current user")
			},
			MockHostname: func() (string, error) {
				return "windows.hostname", nil
			},
			Expected: "Unknown@windows.hostname",
		},
		{
			Desc:     "unable to get hostname",
			UserName: "windows.user",
			MockCurrentUser: func() (*user.User, error) {
				return &user.User{
					Username: "windows.user",
				}, nil
			},
			MockHostname: func() (string, error) {
				return "", errors.New("unable to get hostname")
			},
			Expected: "windows.user@Unknown",
		},
	}

	for _, tc := range tcs {
		getCurrent, getHostname = tc.MockCurrentUser, tc.MockHostname
		t.Run(tc.Desc, func(t *testing.T) {
			windows := New()
			name := windows.GetName()
			if name != tc.Expected {
				t.Fatalf("received %s but expected %s", name, tc.Expected)
			}
		})
	}
}
