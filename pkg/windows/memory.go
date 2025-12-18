package windows

import (
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

var virtualMemory = mem.VirtualMemory

func (w *windows) GetMemoryUsage() string {
	memStat, err := virtualMemory()
	if err != nil {
		return "Unknown"
	}
	total := memStat.Total / (1024 * 1024)
	used := memStat.Used / (1024 * 1024)
	percentage := used * 100 / total
	return fmt.Sprintf("%v MB / %v MB (%v%%)", used, total, percentage)

}
