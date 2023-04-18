package macos

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestTerminalHelper(t *testing.T) {
	defer t.Cleanup(func() {
		os.Stdout = nil
	})
	if os.Getenv("GO_WANT_HELPER_PROCESS_TERM_PROGRAM") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_TERM_PROGRAM_EMPTY") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_FAILURE") != "1" {
		return
	}
	if os.Getenv("GO_WANT_HELPER_PROCESS_TERM_PROGRAM") == "1" {
		fmt.Fprintf(os.Stdout, "Hyper")
		return
	}
	if os.Getenv("GO_WANT_HELPER_PROCESS_TERM_PROGRAM_EMPTY") == "1" {
		fmt.Fprintf(os.Stdout, "")
		return
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
		EnvCommands     []string
		FakeExecCommand func(envs []string) func(command string, arg ...string) *exec.Cmd
	}{
		{
			Desc:     "success - received terminal name",
			Expected: "Hyper",
			EnvCommands: []string{
				"GO_WANT_HELPER_PROCESS_TERM_PROGRAM=1",
			},
			FakeExecCommand: func(envs []string) func(command string, arg ...string) *exec.Cmd {
				return func(command string, args ...string) *exec.Cmd {
					cs := []string{"-test.run=TestTerminalHelper", "--", command}
					cs = append(cs, args...)
					cmd := exec.Command(os.Args[0], cs...)
					cmd.Env = envs
					envs = envs[1:]
					return cmd
				}
			},
		},
		{
			Desc:     "success - received terminal when $TERM_PROGRAM is not bind",
			Expected: "Hyper",
			EnvCommands: []string{
				"GO_WANT_HELPER_PROCESS_TERM_PROGRAM_EMPTY=1",
				"GO_WANT_HELPER_PROCESS_TERM_PROGRAM=1",
			},
			FakeExecCommand: func(envs []string) func(command string, arg ...string) *exec.Cmd {
				return func(command string, args ...string) *exec.Cmd {
					cs := []string{"-test.run=TestTerminalHelper", "--", command}
					cs = append(cs, args...)
					cmd := exec.Command(os.Args[0], cs...)
					cmd.Env = envs
					envs = envs[1:]
					return cmd
				}
			},
		},
		{
			Desc:     "failed - unable to get terminal name",
			Expected: "Unknown",
			EnvCommands: []string{
				"GO_WANT_HELPER_PROCESS_FAILURE=1",
			},
			FakeExecCommand: func(envs []string) func(command string, arg ...string) *exec.Cmd {
				return func(command string, args ...string) *exec.Cmd {
					cs := []string{"-test.run=TestTerminalHelper", "--", command}
					cs = append(cs, args...)
					cmd := exec.Command(os.Args[0], cs...)
					cmd.Env = envs
					envs = envs[1:]
					return cmd
				}
			},
		},
		{
			Desc:     "failed - unable to get terminal name when binding $TERM",
			Expected: "Unknown",
			EnvCommands: []string{
				"GO_WANT_HELPER_PROCESS_TERM_PROGRAM_EMPTY=1",
				"GO_WANT_HELPER_PROCESS_FAILURE=1",
			},
			FakeExecCommand: func(envs []string) func(command string, arg ...string) *exec.Cmd {
				return func(command string, args ...string) *exec.Cmd {
					cs := []string{"-test.run=TestTerminalHelper", "--", command}
					cs = append(cs, args...)
					cmd := exec.Command(os.Args[0], cs...)
					cmd.Env = envs
					envs = envs[1:]
					return cmd
				}
			},
		},
	}
	for _, tt := range tcs {
		t.Run(tt.Desc, func(t *testing.T) {
			execCommand = tt.FakeExecCommand(tt.EnvCommands)
			mac := New()
			terminal := mac.GetTerminalInfo()
			if terminal != tt.Expected {
				t.Fatalf("received %s but expected %s", terminal, tt.Expected)
			}
		})
	}
}
