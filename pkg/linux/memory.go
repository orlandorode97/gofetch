package linux

import (
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

var virtualMemory = mem.VirtualMemory

// GetMemoryUsage returns the memory usage of the computer.
func (l *linux) GetMemoryUsage() string {
	memStat, err := virtualMemory()
	if err != nil {
		return "Unknown"
	}
	total := memStat.Total / (1024 * 1024)
	used := memStat.Used / (1024 * 1024)
	return fmt.Sprintf("%v MB / %v MB", used, total)
}
