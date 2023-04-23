package windows

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestNumberPackagesHelper(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS_SCOOP_PKG_MANAGER") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_CHOCO_PKG_MANAGER") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_PKG_MANAGER_COUNT") != "1" && os.Getenv("GO_WANT_HELPER_PROCESS_PKG_FAILURE") != "1" {
		return
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_SCOOP_PKG_MANAGER") == "1" {
		fmt.Fprintf(os.Stdout, "C:\\ProgramData\\scoop.exe")
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_CHOCO_PKG_MANAGER") == "1" {
		fmt.Fprintf(os.Stdout, "C:\\ProgramData\\choco.exe")
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS_PKG_MANAGER_COUNT") == "1" {
		fmt.Fprintf(os.Stdout, "29")
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
		EnvCommands     []string
		FakeExecCommand func(envs []string) func(command string, arg ...string) *exec.Cmd
	}{
		{
			Desc:     "success - received number of packages of scoop",
			Expected: "29 (scoop)",
			EnvCommands: []string{
				"GO_WANT_HELPER_PROCESS_SCOOP_PKG_MANAGER=1",
				"GO_WANT_HELPER_PROCESS_PKG_MANAGER_COUNT=1",
			},
			FakeExecCommand: func(envs []string) func(command string, arg ...string) *exec.Cmd {
				return func(command string, args ...string) *exec.Cmd {
					cs := []string{"-test.run=TestNumberPackagesHelper", "--", command}
					cs = append(cs, args...)
					cmd := exec.Command(os.Args[0], cs...)
					cmd.Env = envs
					envs = envs[1:] // everytime a command is executed remove the first environment variable from envs to have the next result.
					return cmd
				}
			},
		},
		{
			Desc:     "success - received number of packages of choco",
			Expected: "29 (choco)",
			EnvCommands: []string{
				"GO_WANT_HELPER_PROCESS_PKG_FAILURE=1",
				"GO_WANT_HELPER_PROCESS_CHOCO_PKG_MANAGER=1",
				"GO_WANT_HELPER_PROCESS_PKG_MANAGER_COUNT=1",
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
			Desc:     "failure - unable to get total of packages of scoop",
			Expected: "Unknown",
			EnvCommands: []string{
				"GO_WANT_HELPER_PROCESS_SCOOP_PKG_MANAGER=1",
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
			Desc:     "failure - unable to get total of packages of choco",
			Expected: "Unknown",
			EnvCommands: []string{
				"GO_WANT_HELPER_PROCESS_PKG_FAILURE=1",
				"GO_WANT_HELPER_PROCESS_CHOCO_PKG_MANAGER=1",
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
			Desc:     "failure - unable to get neither scoop nor choco total of packages ",
			Expected: "Unknown",
			EnvCommands: []string{
				"GO_WANT_HELPER_PROCESS_PKG_FAILURE=1",
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
	}

	for _, tt := range tcs {
		t.Run(tt.Desc, func(t *testing.T) {
			execCommand = tt.FakeExecCommand(tt.EnvCommands)
			windows := New()
			total := windows.GetNumberPackages()
			if total != tt.Expected {
				t.Fatalf("received %s but expected %s", total, tt.Expected)
			}
		})
	}
}
