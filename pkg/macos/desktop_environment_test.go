package macos

import (
	"testing"
)

func TestGetShell(t *testing.T) {
	tcs := []struct {
		desc     string
		expected string
	}{
		{
			desc:     "success - received shell name",
			expected: "Aqua",
		},
	}

	for _, tt := range tcs {
		t.Run(tt.desc, func(t *testing.T) {
			macos := New()
			shell := macos.GetDesktopEnvironment()
			if shell != tt.expected {
				t.Fatalf("received %s but expected %s", shell, tt.expected)
			}
		})
	}
}
