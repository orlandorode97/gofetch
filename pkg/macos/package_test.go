package macos

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestPackageHelper(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS_PKG") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_PKG_FAILURE") != "1" {
		return
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_PKG") == "1" {
		fmt.Fprintf(os.Stdout, "        604")
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
			Expected: "604 (brew)",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestPackageHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_PKG=1"}
				return cmd
			},
		},
		{
			Desc:     "failed - unable to get number of packages",
			Expected: "Unknown",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestPackageHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_PKG_FAILURE=1"}
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
			number := mac.GetNumberPackages()
			if number != tc.Expected {
				t.Fatalf("received %s but expected %s", number, tc.Expected)
			}
		})
	}
}
