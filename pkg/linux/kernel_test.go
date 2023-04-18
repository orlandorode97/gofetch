package linux

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestKernelHelper(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS_KERNEL") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_FAILURE") != "1" {
		return
	}
	if os.Getenv("GO_WANT_HELPER_PROCESS_KERNEL") == "1" {
		fmt.Fprintf(os.Stdout, "Linux 5.15.0-40-generic x86_64")
	}
	if os.Getenv("GO_WANT_HELPER_PROCESS_FAILURE") == "1" {
		os.Exit(1)
	}

	os.Exit(0)
}

func TestGetKernelVersion(t *testing.T) {
	tcs := []struct {
		Desc            string
		Expected        string
		FakeExecCommand func(command string, args ...string) *exec.Cmd
	}{
		{
			Desc:     "success - received kernel version",
			Expected: "Linux 5.15.0-40-generic x86_64",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestKernelHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_KERNEL=1"}
				return cmd
			},
		},
		{
			Desc:     "unable to get kernel version",
			Expected: "Unknown",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestKernelHelper", "--", command}
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
			linux := New()
			os := linux.GetKernelVersion()
			if os != tt.Expected {
				t.Fatalf("received %s but expected %s", os, tt.Expected)
			}
		})
	}
}
