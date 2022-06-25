package linux

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestResolutionHelper(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS_RESOLUTION") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_RESOLUTION_FAILURE") != "1" {
		return
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_RESOLUTION") == "1" {
		fmt.Fprintf(os.Stdout, "dimensions:    1024x794 pixels (204x203 millimeters)")
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_RESOLUTION_FAILURE") == "1" {
		os.Exit(1)
	}

	os.Exit(0)
}
func TestGetResolution(t *testing.T) {
	tcs := []struct {
		Desc            string
		Expected        string
		FakeExecCommand func(command string, args ...string) *exec.Cmd
	}{
		{
			Desc:     "success - received resolution dimensions",
			Expected: "1024x794 pixels (204x203 millimeters)",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestResolutionHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_RESOLUTION=1"}
				return cmd
			},
		},
		{
			Desc:     "failure - unable to get resolution dimensions",
			Expected: "Unknown",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestResolutionHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_RESOLUTION_FAILURE=1"}
				return cmd
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Desc, func(t *testing.T) {
			execCommand = tc.FakeExecCommand
			defer func() {
				execCommand = exec.Command
			}()
			linux := New()
			resolution := linux.GetResolution()
			if resolution != tc.Expected {
				t.Fatalf("received %s but expected %s", resolution, tc.Expected)
			}
		})
	}
}
