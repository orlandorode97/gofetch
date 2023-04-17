package windows

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestOSHelper(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS_OS") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_OS_FAILURE") != "1" {
		return
	}
	if os.Getenv("GO_WANT_HELPER_PROCESS_OS") == "1" {
		fmt.Fprintf(os.Stdout, `Caption
Windows 11 Multi Language`)
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
			Desc:     "success - received os",
			Expected: "Windows 11 Multi Language",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestOSHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_OS=1"}
				return cmd
			},
		},
		{
			Desc:     "unable to get os name",
			Expected: "Unknown",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestOSHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_OS_FAILURE=1"}
				return cmd
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Desc, func(t *testing.T) {
			execCommand = tc.FakeExecCommand
			windows := New()
			os := windows.GetOSVersion()
			if os != tc.Expected {
				t.Fatalf("received %s but expected %s", os, tc.Expected)
			}
		})
	}
}
