package stats

import (
	"fmt"
	"runtime"
)

type CPUAndMemoryStats struct {
	CPUUsage      int    `json:"cpu_usage"`
	NumCPUs       int    `json:"number_cpus"`
	AllocMem      uint64 `json:"allocated_memory"`
	TotalAllocMem uint64 `json:"total_allocated_memory"`
	SystemMem     uint64 `json:"system_memory"`
	NumGoRoutines int    `json:"number_go_routines"`
}

func MonitorCPUAndMemory() {

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	stats := CPUAndMemoryStats{
		CPUUsage:      runtime.NumCPU(),
		NumCPUs:       runtime.NumCPU(),
		AllocMem:      memStats.Alloc,
		TotalAllocMem: memStats.TotalAlloc,
		SystemMem:     memStats.Sys,
		NumGoRoutines: runtime.NumGoroutine(),
	}

	fmt.Println(stats)

}
