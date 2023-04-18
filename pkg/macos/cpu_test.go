package macos

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestCPUHelper(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS_CPU") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_FAILURE") != "1" {
		return
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_CPU") == "1" {
		fmt.Fprintf(os.Stdout, "machdep.cpu.brand_string: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz")
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_FAILURE") == "1" {
		os.Exit(1)
	}

	os.Exit(0)
}

func TestGetCPU(t *testing.T) {
	tcs := []struct {
		Desc            string
		Expected        string
		FakeExecCommand func(command string, args ...string) *exec.Cmd
	}{
		{
			Desc:     "success - received cpu name",
			Expected: "Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestCPUHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_CPU=1"}
				return cmd
			},
		},
		{
			Desc:     "failed - unable to get cpu",
			Expected: "Unknown",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestCPUHelper", "--", command}
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
			cpu := mac.GetCPU()
			if cpu != tt.Expected {
				t.Fatalf("received %s but expected %s", cpu, tt.Expected)
			}
		})
	}
}
