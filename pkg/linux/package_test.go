package linux

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

var isCurrentPkgCommand = true

func TestNumberPackagesHelper(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS_PKG_MANAGER") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_PKG_NUMBER") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_PKG_FAILURE") != "1" {
		return
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_PKG_MANAGER") == "1" {
		fmt.Fprintf(os.Stdout, "/usr/bin/pacman")
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_PKG_NUMBER") == "1" {
		fmt.Fprintf(os.Stdout, `warning: database file for 'extra' does not exist
warning: database file for 'community' does not exist
error: failed to prepare transaction (could not find database)
234`)
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_PKG_FAILURE") == "1" {
		os.Exit(1)
	}

	os.Exit(0)
}

func TestGetNumberPackages(t *testing.T) {
	tcs := []struct {
		Desc            string
		Expected        string
		FakeExecCommand func(command string, args ...string) *exec.Cmd
	}{
		{
			Desc:     "success - received number of packages",
			Expected: "234 (pacman)",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestNumberPackagesHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				if isCurrentPkgCommand {
					cmd.Env = []string{"GO_WANT_HELPER_PROCESS_PKG_MANAGER=1"}
					isCurrentPkgCommand = false
					return cmd
				}
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_PKG_NUMBER=1"}
				return cmd
			},
		},
		{
			Desc:     "failed - unable to get current package manager",
			Expected: "Unknown",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestNumberPackagesHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_PKG_FAILURE=1"}
				return cmd
			},
		},
		{
			Desc:     "failed - unable to get the total of the packages",
			Expected: "Unknown",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestNumberPackagesHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				if isCurrentPkgCommand {
					cmd.Env = []string{"GO_WANT_HELPER_PROCESS_PKG_MANAGER=1"}
					isCurrentPkgCommand = false
					return cmd
				}
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_PKG_FAILURE=1"}
				return cmd
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Desc, func(t *testing.T) {
			execCommand = tc.FakeExecCommand

			linux := New()
			total := linux.GetNumberPackages()
			if total != tc.Expected {
				t.Fatalf("received %s but expected %s", total, tc.Expected)
			}

			t.Cleanup(func() {
				isCurrentPkgCommand = true
			})
		})
	}
}
