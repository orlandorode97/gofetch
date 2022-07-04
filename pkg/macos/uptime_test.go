package macos

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

var isBootTimeCommand = true

func TestUptimeHelper(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS_UPTIME_BOOTTIME") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_UPTIME") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_UPTIME_FAILURE") != "1" {
		return
	}
	if os.Getenv("GO_WANT_HELPER_PROCESS_UPTIME_BOOTTIME") == "1" {
		fmt.Fprintf(os.Stdout, "{ sec = 1656944602, usec = 201499 } Mon Jul  4 09:23:22 2022")
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_UPTIME") == "1" {
		fmt.Fprintf(os.Stdout, "123456")
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_UPTIME_FAILURE") == "1" {
		os.Exit(1)
	}

	os.Exit(0)
}
func TestGetUptime(t *testing.T) {
	tcs := []struct {
		Desc            string
		Expected        string
		FakeExecCommand func(command string, args ...string) *exec.Cmd
	}{
		{
			Desc:     "success - received uptime",
			Expected: "1 day(s), 10 hour(s), 17 minutes(s)",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestUptimeHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				if isBootTimeCommand {
					cmd.Env = []string{"GO_WANT_HELPER_PROCESS_UPTIME_BOOTTIME=1"}
					isBootTimeCommand = false
					return cmd
				}
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_UPTIME=1"}
				return cmd
			},
		},
		{
			Desc:     "failure - unable to get boot time",
			Expected: "Unknown",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestUptimeHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_UPTIME_FAILURE=1"}
				return cmd
			},
		},
		{
			Desc:     "failure - unable to get uptime in seconds",
			Expected: "Unknown",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestUptimeHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				if isBootTimeCommand {
					cmd.Env = []string{"GO_WANT_HELPER_PROCESS_UPTIME_BOOTTIME=1"}
					isBootTimeCommand = false
					return cmd
				}
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_UPTIME_FAILURE=1"}
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
			uptime := mac.GetUptime()
			if uptime != tc.Expected {
				t.Fatalf("received %s but expected %s", uptime, tc.Expected)
			}

			t.Cleanup(func() {
				isBootTimeCommand = true
			})
		})
	}
}
