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
		desc            string
		expected        string
		fakeExecCommand func(command string, args ...string) *exec.Cmd
	}{
		{
			desc:     "success - received desktop for windows 10",
			expected: "Fluent",
			fakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestDesktopEnvHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_DE_WINDOWS_10=1"}
				return cmd
			},
		},
		{
			desc:     "success - received desktop environment for windows 8",
			expected: "Metro",
			fakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestDesktopEnvHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_DE_WINDOWS_8=1"}
				return cmd
			},
		},
		{
			desc:     "success - received desktop environment for any windows version",
			expected: "Aero",
			fakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestDesktopEnvHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_DE_WINDOWS_X=1"}
				return cmd
			},
		},
		{
			desc:     "unable to desktop environment name",
			expected: "Unknown",
			fakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestDesktopEnvHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_FAILURE=1"}
				return cmd
			},
		},
	}

	for _, tt := range tcs {
<<<<<<< HEAD:pkg/windows/desktop_environment_test.go
		t.Run(tt.desc, func(t *testing.T) {
			execCommand = tt.fakeExecCommand
			windows := New()
			os := windows.GetDesktopEnvironment()
			if os != tt.expected {
				t.Fatalf("received %s but expected %s", os, tt.expected)
=======
		t.Run(tt.Desc, func(t *testing.T) {
			execCommand = tt.FakeExecCommand
			windows := New()
			os := windows.GetDesktopEnvironment()
			if os != tt.Expected {
				t.Fatalf("received %s but expected %s", os, tt.Expected)
>>>>>>> main:pkg/windows/desk_environment_test.go
			}
		})
	}
}
