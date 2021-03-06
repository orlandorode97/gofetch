package macos

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

var isOSCommand = true

func TestOSHelper(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS_OS_NAME") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_OS_VERSION") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_OS_FAILURE") != "1" {
		return
	}
	if os.Getenv("GO_WANT_HELPER_PROCESS_OS_NAME") == "1" {
		fmt.Fprintf(os.Stdout, "MacOS")
	}
	if os.Getenv("GO_WANT_HELPER_PROCESS_OS_VERSION") == "1" {
		fmt.Fprintf(os.Stdout, "14.01")
	}
	if os.Getenv("GO_WANT_HELPER_PROCESS_OS_FAILURE") == "1" {
		os.Exit(1)
	}
	os.Exit(0)
}

func TestGetOSVersion(t *testing.T) {
	tcs := []struct {
		Desc            string
		Expected        string
		FakeExecCommand func(command string, args ...string) *exec.Cmd
	}{
		{
			Desc:     "success - received os version",
			Expected: "MacOS 14.01",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestOSHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				if isOSCommand {
					cmd.Env = []string{"GO_WANT_HELPER_PROCESS_OS_NAME=1"}
					isOSCommand = false
					return cmd
				}
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_OS_VERSION=1"}
				return cmd
			},
		},
		{
			Desc:     "unable to os name",
			Expected: "Unknown",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestOSHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_OS_FAILURE=1"}
				return cmd
			},
		},
		{
			Desc:     "unable to get os version",
			Expected: "Unknown",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestOSHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				if isOSCommand {
					cmd.Env = []string{"GO_WANT_HELPER_PROCESS_OS_NAME=1"}
					isOSCommand = false
					return cmd
				}
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_OS_FAILURE=1"}
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
			os := linux.GetOSVersion()
			if os != tc.Expected {
				t.Fatalf("received %s but expected %s", os, tc.Expected)
			}

			t.Cleanup(func() {
				isOSCommand = true
			})
		})
	}
}
