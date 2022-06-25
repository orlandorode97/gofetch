package linux

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

const (
	DesktopEnvironment int = iota + 1
	XDG
	XDGVersion
	DEErr
)

var (
	currentXDG = XDG
	currentErr = 0
)

func TestDesktopEnvHelper(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS_DE") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_DE_XDG") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_DE_XDG_VERSION") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_FAILURE") != "1" {
		return
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_DE") == "1" {
		fmt.Fprintf(os.Stdout, "Pantheon")
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_DE_XDG") == "1" {
		fmt.Fprintf(os.Stdout, "ubuntu:GNOME")
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_DE_XDG_VERSION") == "1" {
		fmt.Fprintf(os.Stdout, "GNOME Shell 40.2")
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
			Desc:     "success - received desktop environment",
			Expected: "Pantheon",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestDesktopEnvHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_DE=1"}
				return cmd
			},
		},
		{
			Desc:     "success - received desktop environment with $XDG extended",
			Expected: "GNOME Shell 40.2",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestDesktopEnvHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				if currentXDG == XDG {
					cmd.Env = []string{"GO_WANT_HELPER_PROCESS_DE_XDG=1"}
					currentXDG = XDGVersion
					return cmd
				}
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_DE_XDG_VERSION=1"}
				return cmd
			},
		},
		{
			Desc:     "unable to $XDG_CURRENT_DESKTOP",
			Expected: "Unknown",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestDesktopEnvHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_FAILURE=1"}
				return cmd
			},
		},
		{
			Desc:     "unable to desktop environment version",
			Expected: "Unknown",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestDesktopEnvHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				if currentXDG == XDG && currentErr == 0 {
					cmd.Env = []string{"GO_WANT_HELPER_PROCESS_DE_XDG=1"}
					currentErr = XDGVersion
					return cmd
				}
				if currentErr == XDGVersion {
					cmd.Env = []string{"GO_WANT_HELPER_PROCESS_FAILURE=1"}
					return cmd
				}
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
			linux := New()
			os := linux.GetDesktopEnvironment()
			if os != tc.Expected {
				t.Fatalf("received %s but expected %s", os, tc.Expected)
			}

			t.Cleanup(func() {
				currentXDG = XDG
				currentErr = 0
			})
		})
	}
}
