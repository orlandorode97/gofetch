package macos

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

var isTerminalCommand = true

func TestTerminalHelper(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS_TERM_PROGRAM") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_TERM_PROGRAM_EMPTY") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_TERMINAL_FAILURE") != "1" {
		return
	}
	if os.Getenv("GO_WANT_HELPER_PROCESS_TERM_PROGRAM") == "1" {
		fmt.Fprintf(os.Stdout, "Hyper")
	}
	if os.Getenv("GO_WANT_HELPER_PROCESS_TERM_PROGRAM_EMPTY") == "1" {
		fmt.Fprintf(os.Stdout, "")
	}
	if os.Getenv("GO_WANT_HELPER_PROCESS_TERMINAL_FAILURE") == "1" {
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
			Expected: "Hyper",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestTerminalHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_TERM_PROGRAM=1"}
				return cmd
			},
		},
		{
			Desc:     "success - received terminal when $TERM_PROGRAM is not bind",
			Expected: "Hyper",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestTerminalHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				if isTerminalCommand {
					cmd.Env = []string{"GO_WANT_HELPER_PROCESS_TERM_PROGRAM_EMPTY=1"}
					isTerminalCommand = false
					return cmd
				}
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_TERM_PROGRAM=1"}
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
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_TERMINAL_FAILURE=1"}
				return cmd
			},
		},
		{
			Desc:     "failed - unable to get terminal name when binding $TERM",
			Expected: "Unknown",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestTerminalHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				if isTerminalCommand {
					cmd.Env = []string{"GO_WANT_HELPER_PROCESS_TERM_PROGRAM_EMPTY=1"}
					isTerminalCommand = false
					return cmd
				}
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_TERMINAL_FAILURE=1"}
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
			terminal := mac.GetTerminalInfo()
			if terminal != tc.Expected {
				t.Fatalf("received %s but expected %s", terminal, tc.Expected)
			}

			t.Cleanup(func() {
				isTerminalCommand = true
			})
		})
	}
}
