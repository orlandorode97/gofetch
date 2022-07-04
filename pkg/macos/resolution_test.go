package macos

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
		fmt.Fprintf(os.Stdout, `Resolution: 3072 x 1920 Retina
          Resolution: 1920 x 1080 (1080p FHD - Full High Definition)`)
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
			Desc:     "success - received mac resolution",
			Expected: "3072 x 1920 Retina",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestResolutionHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_RESOLUTION=1"}
				return cmd
			},
		},
		{
			Desc:     "failed - unable to get resolution",
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
			mac := New()
			cpu := mac.GetResolution()
			if cpu != tc.Expected {
				t.Fatalf("received %s but expected %s", cpu, tc.Expected)
			}
		})
	}
}
