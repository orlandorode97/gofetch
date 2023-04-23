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

	for _, tt := range tcs {
		t.Run(tt.Desc, func(t *testing.T) {
			windows := New()
			shell := windows.GetShellInformation()
			if shell != tt.Expected {
				t.Fatalf("received %s but expected %s", shell, tt.Expected)
			}
		})
	}
}
