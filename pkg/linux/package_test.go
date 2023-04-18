package linux

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestNumberPackagesHelper(t *testing.T) {
	defer t.Cleanup(func() { // Clean standard output since after running this helper it ends with "PASS"
		os.Stdout = nil
	})

	if os.Getenv("GO_WANT_HELPER_PROCESS_PKG_MANAGER") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_PKG_NUMBER") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_FAILURE") != "1" {
		return
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_PKG_MANAGER") == "1" {
		fmt.Fprintf(os.Stderr, "/usr/bin/pacman")
		return
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_PKG_NUMBER") == "1" {
		fmt.Fprintf(os.Stderr, `warning: database file for 'extra' does not exist
warning: database file for 'community' does not exist
error: failed to prepare transaction (could not find database)
234`)
		return
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_FAILURE") == "1" {
		os.Exit(1)
		return
	}

	os.Exit(0)
}

func TestGetNumberPackages(t *testing.T) {
	tcs := []struct {
		Desc            string
		Expected        string
		EnvCommands     []string
		FakeExecCommand func(envs []string) func(command string, arg ...string) *exec.Cmd
	}{
		{
			Desc:     "success - received number of packages",
			Expected: "234 (pacman)",
			EnvCommands: []string{
				"GO_WANT_HELPER_PROCESS_PKG_MANAGER=1",
				"GO_WANT_HELPER_PROCESS_PKG_NUMBER=1",
			},
			FakeExecCommand: func(envs []string) func(command string, arg ...string) *exec.Cmd {
				return func(command string, args ...string) *exec.Cmd {
					cs := []string{"-test.run=TestNumberPackagesHelper", "--", command}
					cs = append(cs, args...)
					cmd := exec.Command(os.Args[0], cs...)
					cmd.Env = envs
					envs = envs[1:]
					return cmd
				}
			},
		},
		{
			Desc:     "failed - unable to get current package manager",
			Expected: "Unknown",
			EnvCommands: []string{
				"GO_WANT_HELPER_PROCESS_PKG_FAILURE=1",
			},
			FakeExecCommand: func(envs []string) func(command string, arg ...string) *exec.Cmd {
				return func(command string, args ...string) *exec.Cmd {
					cs := []string{"-test.run=TestNumberPackagesHelper", "--", command}
					cs = append(cs, args...)
					cmd := exec.Command(os.Args[0], cs...)
					cmd.Env = envs
					envs = envs[1:]
					return cmd
				}
			},
		},
		{
			Desc:     "failed - unable to get the total of the packages",
			Expected: "Unknown",
			EnvCommands: []string{
				"GO_WANT_HELPER_PROCESS_PKG_MANAGER=1",
				"GO_WANT_HELPER_PROCESS_FAILURE=1",
			},
			FakeExecCommand: func(envs []string) func(command string, arg ...string) *exec.Cmd {
				return func(command string, args ...string) *exec.Cmd {
					cs := []string{"-test.run=TestNumberPackagesHelper", "--", command}
					cs = append(cs, args...)
					cmd := exec.Command(os.Args[0], cs...)
					cmd.Env = envs
					envs = envs[1:]
					return cmd
				}
			},
		},
	}

	for _, tt := range tcs {
		t.Run(tt.Desc, func(t *testing.T) {
			execCommand = tt.FakeExecCommand(tt.EnvCommands)
			linux := New()
			total := linux.GetNumberPackages()
			if total != tt.Expected {
				t.Fatalf("received %s but expected %s", total, tt.Expected)
			}
		})
	}
}
