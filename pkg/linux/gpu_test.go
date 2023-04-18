package linux

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestGPUHelper(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS_GPU_VGA") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_GPU_3D") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_GPU_DISPLAY") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_FAILURE") != "1" {
		return
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_GPU_VGA") == "1" {
		fmt.Fprintf(os.Stdout, "00:0f.0 VGA compatible controller: VMware SVGA II Adapter")
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_GPU_3D") == "1" {
		fmt.Fprintf(os.Stdout, "01:00.0 3D controller: NVIDIA Corporation GK107M [GeForce GT 750M] (rev a1)")
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_GPU_DISPLAY") == "1" {
		fmt.Fprintf(os.Stdout, "00:0e.1 Display controller: AMD Radeon RX 5600 XT (rev 01)")
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_FAILURE") == "1" {
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
			Desc:     "success - received vga gpu name",
			Expected: "VMware SVGA II Adapter",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestGPUHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_GPU_VGA=1"}
				return cmd
			},
		},
		{
			Desc:     "success - received 3D gpu name",
			Expected: "NVIDIA Corporation GK107M [GeForce GT 750M]",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestGPUHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_GPU_3D=1"}
				return cmd
			},
		},
		{
			Desc:     "success - received display gpu name",
			Expected: "AMD Radeon RX 5600 XT",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestGPUHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_GPU_DISPLAY=1"}
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
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_FAILURE=1"}
				return cmd
			},
		},
	}

	for _, tt := range tcs {
		t.Run(tt.Desc, func(t *testing.T) {
			execCommand = tt.FakeExecCommand
			linux := New()
			gpu := linux.GetGPU()
			if gpu != tt.Expected {
				t.Fatalf("received %s but expected %s", gpu, tt.Expected)
			}
		})
	}
}
