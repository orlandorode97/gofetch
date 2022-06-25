package linux

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
			Desc:     "success - received linux name",
			UserName: "linux.user",
			Hostname: "linux.hostname",
			MockCurrentUser: func() (*user.User, error) {
				return &user.User{
					Username: "linux.user",
				}, nil
			},
			MockHostname: func() (string, error) {
				return "linux.hostname", nil
			},
			Expected: "linux.user@linux.hostname",
		},
		{
			Desc: "unable to get current user",
			MockCurrentUser: func() (*user.User, error) {
				return &user.User{}, errors.New("unable to get current user")
			},
			MockHostname: func() (string, error) {
				return "linux.hostname", nil
			},
			Expected: "Unknown@linux.hostname",
		},
		{
			Desc:     "unable to get hostname",
			UserName: "linux.user",
			MockCurrentUser: func() (*user.User, error) {
				return &user.User{
					Username: "linux.user",
				}, nil
			},
			MockHostname: func() (string, error) {
				return "", errors.New("unable to get hostname")
			},
			Expected: "linux.user@Unknown",
		},
	}

	for _, tc := range tcs {
		getCurrent, getHostname = tc.MockCurrentUser, tc.MockHostname
		t.Run(tc.Desc, func(t *testing.T) {
			linux := New()
			name := linux.GetName()
			if name != tc.Expected {
				t.Fatalf("received %s but expected %s", name, tc.Expected)
			}
		})
	}
}
