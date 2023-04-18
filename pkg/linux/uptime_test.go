package linux

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestUptimeHelper(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS_UPTIME") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_FAILURE") != "1" {
		return
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_UPTIME") == "1" {
		fmt.Fprintf(os.Stdout, "123456")
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_FAILURE") == "1" {
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
			Expected: "1 day(s), 10 hour(s), 17 minute(s)",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestUptimeHelper", "--", command}
				cs = append(cs, args...)
				cmd := exec.Command(os.Args[0], cs...)
				cmd.Env = []string{"GO_WANT_HELPER_PROCESS_UPTIME=1"}
				return cmd
			},
		},
		{
			Desc:     "failure - unable to get uptime",
			Expected: "Unknown",
			FakeExecCommand: func(command string, args ...string) *exec.Cmd {
				cs := []string{"-test.run=TestUptimeHelper", "--", command}
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
			defer func() {
				execCommand = exec.Command
			}()
			linux := New()
			uptime := linux.GetUptime()
			if uptime != tt.Expected {
				t.Fatalf("received %s but expected %s", uptime, tt.Expected)
			}
		})
	}
}
