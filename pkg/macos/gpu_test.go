package macos

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestGPUHelper(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS_GPU") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_GPU_FAILURE") != "1" {
		return
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_GPU") == "1" {
		fmt.Fprintf(os.Stdout, `
      Chipset Model: Intel UHD Graphics 630
      Chipset Model: AMD Radeon Pro 5300M`)
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_GPU_FAILURE") == "1" {
		os.Exit(1)
	}

	os.Exit(0)
}

func TestGetGPU(t *testing.T) {
	tcs := []struct {
		Desc            string
		Expected        string
		FakeExecCommand func(command string, args ...string) *exec.Cmd
	}{
		{
			Desc:     "success - received gpu name",
			Expected: "Intel UHD Graphics 630",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestGPUHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_GPU=1"}
				return cmd
			},
		},
		{
			Desc:     "failed - unable to get gpu name",
			Expected: "Unknown",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestGPUHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_GPU_FAILURE=1"}
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
			gpu := mac.GetGPU()
			if gpu != tc.Expected {
				t.Fatalf("received %s but expected %s", gpu, tc.Expected)
			}
		})
	}
}
