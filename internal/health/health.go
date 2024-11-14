package health

import (
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

// FetchMetrics fetches real system metrics for CPU and memory usage
func FetchMetrics() {
	cpuUsage, _ := cpu.Percent(0, false)
	vmStat, _ := mem.VirtualMemory()

	CPUUsageGauge.Set(cpuUsage[0])
	MemoryUsageGauge.Set(vmStat.UsedPercent)
}

func FetchMockMetrics() {
	mockCPU := 25.0
	mockMemory := 60.0

	CPUUsageGauge.Set(mockCPU)
	MemoryUsageGauge.Set(mockMemory)
}
