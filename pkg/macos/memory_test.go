package macos

import (
	"errors"
	"testing"

	"github.com/shirou/gopsutil/mem"
)

func TestGetMemoryUsage(t *testing.T) {
	tcs := []struct {
		Desc          string
		Expected      string
		VirtualMemory func() (*mem.VirtualMemoryStat, error)
	}{
		{
			Desc:     "success - received memory usage",
			Expected: "6656 MB / 16384 MB",
			VirtualMemory: func() (*mem.VirtualMemoryStat, error) {
				return &mem.VirtualMemoryStat{
					Total: 0x400000000, // 16 GB
					Used:  0x1a0000000, // 6.5 GB
				}, nil
			},
		},
		{
			Desc:     "failed - unable to get memory usage",
			Expected: "Unknown",
			VirtualMemory: func() (*mem.VirtualMemoryStat, error) {
				return nil, errors.New("unable to get memory stats")
			},
		},
	}

	for _, tt := range tcs {
		t.Run(tt.Desc, func(t *testing.T) {
			virtualMemory = tt.VirtualMemory
			mac := New()
			memory := mac.GetMemoryUsage()
			if memory != tt.Expected {
				t.Fatalf("received %s but expected %s", memory, tt.Expected)
			}
		})
	}
}
