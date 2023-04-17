package windows

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestCPUHelper(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS_CPU") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_CPU_FAILURE") != "1" {
		return
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_CPU") == "1" {
		fmt.Fprintf(os.Stdout, `Name
Intel(R) Core(TM) i5-1020U`)
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_CPU_FAILURE") == "1" {
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
			Expected: "Intel(R) Core(TM) i5-1020U",
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
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_CPU_FAILURE=1"}
				return cmd
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Desc, func(t *testing.T) {
			execCommand = tc.FakeExecCommand
			windows := New()
			cpu := windows.GetCPU()
			if cpu != tc.Expected {
				t.Fatalf("received %s but expected %s", cpu, tc.Expected)
			}
		})
	}
}
