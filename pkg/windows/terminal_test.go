package windows

import (
	"testing"
)

func TestTerminal(t *testing.T) {
	tcs := []struct {
		Desc     string
		Expected string
	}{
		{
			Desc:     "success - received terminal name",
			Expected: "Windows terminal",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Desc, func(t *testing.T) {
			windows := New()
			terminal := windows.GetTerminalInfo()
			if terminal != tc.Expected {
				t.Fatalf("received %s but expected %s", terminal, tc.Expected)
			}
		})
	}
}
