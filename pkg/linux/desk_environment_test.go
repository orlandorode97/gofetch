package linux

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestDesktopEnvHelper(t *testing.T) {
	defer t.Cleanup(func() {
		os.Stdout = nil
	})
	if os.Getenv("GO_WANT_HELPER_PROCESS_DE") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_DE_XDG") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_DE_XDG_VERSION") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_FAILURE") != "1" {
		return
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_DE") == "1" {
		fmt.Fprintf(os.Stderr, "Pantheon")
		return
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_DE_XDG") == "1" {
		fmt.Fprintf(os.Stderr, "ubuntu:GNOME")
		return
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_DE_XDG_VERSION") == "1" {
		fmt.Fprintf(os.Stderr, "GNOME Shell 40.2")
		return
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_FAILURE") == "1" {
		os.Exit(1)
		return
	}

	os.Exit(0)
}

func TestGetDesktopEnvironment(t *testing.T) {
	tcs := []struct {
		Desc            string
		Expected        string
		EnvCommands     []string
		FakeExecCommand func(envs []string) func(command string, arg ...string) *exec.Cmd
	}{
		{
			Desc:        "success - received desktop environment",
			Expected:    "Pantheon",
			EnvCommands: []string{"GO_WANT_HELPER_PROCESS_DE=1"},
			FakeExecCommand: func(envs []string) func(command string, arg ...string) *exec.Cmd {
				return func(command string, args ...string) *exec.Cmd {
					cs := []string{"-test.run=TestDesktopEnvHelper", "--", command}
					cs = append(cs, args...)
					cmd := exec.Command(os.Args[0], cs...)
					cmd.Env = envs
					envs = envs[1:] // everytime a command is executed remove the first environment variable from envs to have the next result.
					return cmd
				}
			},
		},
		{
			Desc:     "success - received desktop environment with $XDG extended",
			Expected: "GNOME Shell 40.2",
			EnvCommands: []string{
				"GO_WANT_HELPER_PROCESS_DE_XDG=1",
				"GO_WANT_HELPER_PROCESS_DE_XDG_VERSION=1",
			},
			FakeExecCommand: func(envs []string) func(command string, arg ...string) *exec.Cmd {
				return func(command string, args ...string) *exec.Cmd {
					cs := []string{"-test.run=TestDesktopEnvHelper", "--", command}
					cs = append(cs, args...)
					cmd := exec.Command(os.Args[0], cs...)
					cmd.Env = envs
					envs = envs[1:] // everytime a command is executed remove the first environment variable from envs to have the next result.
					return cmd
				}
			},
		},
		{
			Desc:     "failure - unable to $XDG_CURRENT_DESKTOP",
			Expected: "Unknown",
			EnvCommands: []string{
				"GO_WANT_HELPER_PROCESS_FAILURE=1",
			},
			FakeExecCommand: func(envs []string) func(command string, arg ...string) *exec.Cmd {
				return func(command string, args ...string) *exec.Cmd {
					cs := []string{"-test.run=TestDesktopEnvHelper", "--", command}
					cs = append(cs, args...)
					cmd := exec.Command(os.Args[0], cs...)
					cmd.Env = envs
					envs = envs[1:] // everytime a command is executed remove the first environment variable from envs to have the next result.
					return cmd
				}
			},
		},
		{
			Desc:     "failure - unable to desktop environment version",
			Expected: "Unknown",
			EnvCommands: []string{
				"GO_WANT_HELPER_PROCESS_DE_XDG=1",
				"GO_WANT_HELPER_PROCESS_FAILURE=1",
			},
			FakeExecCommand: func(envs []string) func(command string, arg ...string) *exec.Cmd {
				return func(command string, args ...string) *exec.Cmd {
					cs := []string{"-test.run=TestDesktopEnvHelper", "--", command}
					cs = append(cs, args...)
					cmd := exec.Command(os.Args[0], cs...)
					cmd.Env = envs
					envs = envs[1:] // everytime a command is executed remove the first environment variable from envs to have the next result.
					return cmd
				}
			},
		},
	}

	for _, tt := range tcs {
		t.Run(tt.Desc, func(t *testing.T) {
			execCommand = tt.FakeExecCommand(tt.EnvCommands)
			linux := New()
			os := linux.GetDesktopEnvironment()
			if os != tt.Expected {
				t.Fatalf("received %s but expected %s", os, tt.Expected)
			}
		})
	}
}
