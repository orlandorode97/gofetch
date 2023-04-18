package macos

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestOSHelper(t *testing.T) {
	defer t.Cleanup(func() {
		os.Stdout = nil
	})
	if os.Getenv("GO_WANT_HELPER_PROCESS_OS_NAME") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_OS_VERSION") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_FAILURE") != "1" {
		return
	}
	if os.Getenv("GO_WANT_HELPER_PROCESS_OS_NAME") == "1" {
		fmt.Fprintf(os.Stdout, "MacOS")
		return
	}
	if os.Getenv("GO_WANT_HELPER_PROCESS_OS_VERSION") == "1" {
		fmt.Fprintf(os.Stdout, "14.01")
		return
	}
	if os.Getenv("GO_WANT_HELPER_PROCESS_FAILURE") == "1" {
		os.Exit(1)
	}
	os.Exit(0)
}

func TestGetOSVersion(t *testing.T) {
	tcs := []struct {
		Desc            string
		Expected        string
		EnvCommands     []string
		FakeExecCommand func(envs []string) func(command string, arg ...string) *exec.Cmd
	}{
		{
			Desc:     "success - received os version",
			Expected: "MacOS 14.01",
			EnvCommands: []string{
				"GO_WANT_HELPER_PROCESS_OS_NAME=1",
				"GO_WANT_HELPER_PROCESS_OS_VERSION=1",
			},
			FakeExecCommand: func(envs []string) func(command string, arg ...string) *exec.Cmd {
				return func(command string, args ...string) *exec.Cmd {
					cs := []string{"-test.run=TestOSHelper", "--", command}
					cs = append(cs, args...)
					cmd := exec.Command(os.Args[0], cs...)
					cmd.Env = envs
					envs = envs[1:]
					return cmd
				}
			},
		},
		{
			Desc:     "failure - unable to os name",
			Expected: "Unknown",
			EnvCommands: []string{
				"GO_WANT_HELPER_PROCESS_FAILURE=1",
			},
			FakeExecCommand: func(envs []string) func(command string, arg ...string) *exec.Cmd {
				return func(command string, args ...string) *exec.Cmd {
					cs := []string{"-test.run=TestOSHelper", "--", command}
					cs = append(cs, args...)
					cmd := exec.Command(os.Args[0], cs...)
					cmd.Env = envs
					envs = envs[1:]
					return cmd
				}
			},
		},
		{
			Desc:     "unable to get os version",
			Expected: "Unknown",
			EnvCommands: []string{
				"GO_WANT_HELPER_PROCESS_OS_NAME=1",
				"GO_WANT_HELPER_PROCESS_FAILURE=1",
			},
			FakeExecCommand: func(envs []string) func(command string, arg ...string) *exec.Cmd {
				return func(command string, args ...string) *exec.Cmd {
					cs := []string{"-test.run=TestOSHelper", "--", command}
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
			linux := New()
			os := linux.GetOSVersion()
			if os != tt.Expected {
				t.Fatalf("received %s but expected %s", os, tt.Expected)
			}
		})
	}
}
