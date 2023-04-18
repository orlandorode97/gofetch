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

	for _, tt := range tcs {
		t.Run(tt.Desc, func(t *testing.T) {
			windows := New()
			terminal := windows.GetTerminalInfo()
			if terminal != tt.Expected {
				t.Fatalf("received %s but expected %s", terminal, tt.Expected)
			}
		})
	}
}
