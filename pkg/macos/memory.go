package macos

import (
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

// GetMemoryUsage returns the memory usage of the computer.
func (m *macos) GetMemoryUsage() string {
	memStat, err := mem.VirtualMemory()
	if err != nil {
		return "Unknown"
	}
	total := memStat.Total / (1024 * 1024)
	used := memStat.Used / (1024 * 1024)
	return fmt.Sprintf("%v MB / %v MB", used, total)
}
