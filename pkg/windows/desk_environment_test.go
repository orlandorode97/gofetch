package windows

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

var isXDGCommand = true

func TestDesktopEnvHelper(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS_DE_WINDOWS_X") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_DE_WINDOWS_10") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_DE_WINDOWS_8") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_FAILURE") != "1" {
		return
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_DE_WINDOWS_X") == "1" {
		fmt.Fprintf(os.Stdout, "Windows 11 Language Multi")
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_DE_WINDOWS_8") == "1" {
		fmt.Fprintf(os.Stdout, "Windows 8 Language Multi")
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_DE_WINDOWS_10") == "1" {
		fmt.Fprintf(os.Stdout, "Windows 10 Language Multi")
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_FAILURE") == "1" {
		os.Exit(1)
	}

	os.Exit(0)
}

func TestGetDesktopEnvironment(t *testing.T) {
	tcs := []struct {
		Desc            string
		Expected        string
		FakeExecCommand func(command string, args ...string) *exec.Cmd
	}{
		{
			Desc:     "success - received desktop for windows 10",
			Expected: "Fluent",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestDesktopEnvHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_DE_WINDOWS_10=1"}
				return cmd
			},
		},
		{
			Desc:     "success - received desktop environment for windows 8",
			Expected: "Metro",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestDesktopEnvHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_DE_WINDOWS_8=1"}
				return cmd
			},
		},
		{
			Desc:     "success - received desktop environment for any windows version",
			Expected: "Aero",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestDesktopEnvHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_DE_WINDOWS_X=1"}
				return cmd
			},
		},
		{
			Desc:     "unable to desktop environment name",
			Expected: "Unknown",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestDesktopEnvHelper", "--", command}
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
			windows := New()
			os := windows.GetDesktopEnvironment()
			if os != tt.Expected {
				t.Fatalf("received %s but expected %s", os, tt.Expected)
			}
		})
	}
}
