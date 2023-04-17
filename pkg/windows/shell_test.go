package windows

import (
	"testing"
)

func TestGetShell(t *testing.T) {
	tcs := []struct {
		Desc     string
		Expected string
	}{
		{
			Desc:     "success - received shell name",
			Expected: "Unknown",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Desc, func(t *testing.T) {
			windows := New()
			shell := windows.GetShellInformation()
			if shell != tc.Expected {
				t.Fatalf("received %s but expected %s", shell, tc.Expected)
			}
		})
	}
}
