package macos

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestShellHelper(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS_SHELL") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_FAILURE") != "1" {
		return
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_SHELL") == "1" {
		fmt.Fprintf(os.Stdout, "/usr/bin/zsh 5.19.4 (x86_darwin)")
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_FAILURE") == "1" {
		os.Exit(1)
	}

	os.Exit(0)
}
func TestGetShellInformation(t *testing.T) {
	tcs := []struct {
		Desc            string
		Expected        string
		FakeExecCommand func(command string, args ...string) *exec.Cmd
	}{
		{
			Desc:     "success - received shell information",
			Expected: "zsh 5.19.4",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestShellHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_SHELL=1"}
				return cmd
			},
		},
		{
			Desc:     "failure - unable to get shell information",
			Expected: "Unknown",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestShellHelper", "--", command}
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
			mac := New()
			uptime := mac.GetShellInformation()
			if uptime != tt.Expected {
				t.Fatalf("received %s but expected %s", uptime, tt.Expected)
			}
		})
	}
}
