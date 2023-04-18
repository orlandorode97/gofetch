package linux

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestTerminalHelper(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS_TERMINAL") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_FAILURE") != "1" {
		return
	}
	if os.Getenv("GO_WANT_HELPER_PROCESS_TERMINAL") == "1" {
		fmt.Fprintf(os.Stdout, "Alacritty")
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_FAILURE") == "1" {
		os.Exit(1)
	}
	os.Exit(0)
}

func TestGetTerminalInfo(t *testing.T) {
	tcs := []struct {
		Desc            string
		Expected        string
		FakeExecCommand func(command string, args ...string) *exec.Cmd
	}{
		{
			Desc:     "success - received terminal name",
			Expected: "Alacritty",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestTerminalHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_TERMINAL=1"}
				return cmd
			},
		},
		{
			Desc:     "failed - unable to get terminal name",
			Expected: "Unknown",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestTerminalHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_FAILURE=1"}
				return cmd
			},
		},
	}
	for _, tt := range tcs {
		t.Run(tt.Desc, func(t *testing.T) {
			execCommand = tt.FakeExecCommand
			defer func() {
				execCommand = exec.Command
			}()
			linux := New()
			terminal := linux.GetTerminalInfo()
			if terminal != tt.Expected {
				t.Fatalf("received %s but expected %s", terminal, tt.Expected)
			}
		})
	}
}
