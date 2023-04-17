package windows

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestGetResolutionHelper(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS_HORIZONTAL") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_VERTICAL") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_FAILURE") != "1" {
		return
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_HORIZONTAL") == "1" {
		fmt.Fprintf(os.Stdout, `CurrentHorizontalResolution
1920`)
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_VERTICAL") == "1" {
		fmt.Fprintf(os.Stdout, `CurrentVerticalResolution
1080`)
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_FAILURE") == "1" {
		os.Exit(1)
	}

	os.Exit(0)
}

func TestGetResolution(t *testing.T) {
	tcs := []struct {
		Desc            string
		Expected        string
		EnvCommands     []string
		FakeExecCommand func(envs []string) func(command string, arg ...string) *exec.Cmd
	}{
		{
			Desc:     "success - received screen resolution",
			Expected: "1920 X 1080",
			EnvCommands: []string{
				"GO_WANT_HELPER_PROCESS_HORIZONTAL=1",
				"GO_WANT_HELPER_PROCESS_VERTICAL=1",
			},
			FakeExecCommand: func(envs []string) func(command string, arg ...string) *exec.Cmd {
				return func(command string, args ...string) *exec.Cmd {
					cs := []string{"-test.run=TestGetResolutionHelper", "--", command}
					cs = append(cs, args...)
					cmd := exec.Command(os.Args[0], cs...)
					cmd.Env = envs
					envs = envs[1:] // everytime a command is executed remove the first environment variable from envs to have the next result.
					return cmd
				}
			},
		},
		{
			Desc:     "failure - unable to get horizontal resolution value",
			Expected: "Unknown",
			EnvCommands: []string{
				"GO_WANT_HELPER_PROCESS_FAILURE=1",
			},
			FakeExecCommand: func(envs []string) func(command string, arg ...string) *exec.Cmd {
				return func(command string, args ...string) *exec.Cmd {
					cs := []string{"-test.run=TestGetResolutionHelper", "--", command}
					cs = append(cs, args...)
					cmd := exec.Command(os.Args[0], cs...)
					cmd.Env = envs
					envs = envs[1:]
					return cmd
				}
			},
		},
		{
			Desc:     "failure - unable to get vertical resolution value",
			Expected: "Unknown",
			EnvCommands: []string{
				"GO_WANT_HELPER_PROCESS_HORIZONTAL",
				"GO_WANT_HELPER_PROCESS_FAILURE=1",
			},
			FakeExecCommand: func(envs []string) func(command string, arg ...string) *exec.Cmd {
				return func(command string, args ...string) *exec.Cmd {
					cs := []string{"-test.run=TestGetResolutionHelper", "--", command}
					cs = append(cs, args...)
					cmd := exec.Command(os.Args[0], cs...)
					cmd.Env = envs
					envs = envs[1:]
					return cmd
				}
			},
		},
		{
			Desc:     "failure - unable to get screen resolution",
			Expected: "Unknown",
			EnvCommands: []string{
				"GO_WANT_HELPER_PROCESS_FAILURE=1",
				"GO_WANT_HELPER_PROCESS_FAILURE=1",
			},
			FakeExecCommand: func(envs []string) func(command string, arg ...string) *exec.Cmd {
				return func(command string, args ...string) *exec.Cmd {
					cs := []string{"-test.run=TestGetResolutionHelper", "--", command}
					cs = append(cs, args...)
					cmd := exec.Command(os.Args[0], cs...)
					cmd.Env = envs
					envs = envs[1:]
					return cmd
				}
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Desc, func(t *testing.T) {
			execCommand = tc.FakeExecCommand(tc.EnvCommands)
			windows := New()
			resolution := windows.GetResolution()
			if resolution != tc.Expected {
				t.Fatalf("received %s but expected %s", resolution, tc.Expected)
			}
		})
	}
}
