package macos

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestKernelHelper(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCCES_KERNEL") != "1" && os.Getenv("GO_WANT_HELPER_PROCCES_KERNEL_FAILURE") != "1" {
		return
	}

	if os.Getenv("GO_WANT_HELPER_PROCCES_KERNEL") == "1" {
		fmt.Fprintf(os.Stdout, "Kernel version 21.5.0: Tue Apr 26 21:08:22 PDT 2022; root:xnu-8020.121.3~4/RELEASE_X86_64 ")
	}

	if os.Getenv("GO_WANT_HELPER_PROCCES_KERNEL_FAILURE") == "1" {
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
			Expected: "Kernel version 21.5.0",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestKernelHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				cmd.Env = []string{"GO_WANT_HELPER_PROCCES_KERNEL=1"}
				return cmd
			},
		},
		{
			Desc:     "failed - unable to get kernel version",
			Expected: "Unknown",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestKernelHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				cmd.Env = []string{"GO_WANT_HELPER_PROCCES_KERNEL_FAILURE=1"}
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
			cpu := mac.GetKernelVersion()
			if cpu != tc.Expected {
				t.Fatalf("received %s but expected %s", cpu, tc.Expected)
			}
		})
	}
}
