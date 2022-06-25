package linux

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

const (
	pkgManager int = iota + 1
	pkgTotal
	pkgErr
)

var (
	currentPkgCommand = pkgManager
	currentPkgErr     = 0
)

func TestNumberPackagesHelper(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS_PKG_MANAGER") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_PKG_NUMBER") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_PKG_FAILURE") != "1" {
		return
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_PKG_MANAGER") == "1" {
		fmt.Fprintf(os.Stdout, "/usr/bin/pacman")
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_PKG_NUMBER") == "1" {
		fmt.Fprintf(os.Stdout, "234")
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
			Expected: "234",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestNumberPackagesHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				if currentPkgCommand == pkgManager {
					cmd.Env = []string{"GO_WANT_HELPER_PROCESS_PKG_MANAGER=1"}
					currentPkgCommand = pkgTotal
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
				if currentPkgCommand == pkgManager && currentPkgErr == 0 {
					cmd.Env = []string{"GO_WANT_HELPER_PROCESS_PKG_MANAGER=1"}
					currentPkgErr = pkgErr
					return cmd
				}
				if currentPkgErr == pkgErr {
					cmd.Env = []string{"GO_WANT_HELPER_PROCESS_PKG_FAILURE=1"}
				}
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
				currentPkgCommand = pkgManager
				currentPkgErr = 0
			})
		})
	}
}
