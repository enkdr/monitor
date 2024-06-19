package stats

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

type CPUAndMemoryStats struct {
	CPUUsage      int       `json:"cpu_usage"`
	NumCPUs       int       `json:"number_cpus"`
	AllocMem      uint64    `json:"allocated_memory"`
	TotalAllocMem uint64    `json:"total_allocated_memory"`
	SystemMem     uint64    `json:"system_memory"`
	NumGoRoutines int       `json:"number_go_routines"`
	CreatedAt     time.Time `json:"created_at"`
}

func MonitorCPUAndMemory(dbFlag bool) {

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

	jsonData, err := StatsJSON(stats)

	if err != nil {
		fmt.Println(err)
	}

	if dbFlag {
		err = StatsDBInsert("cpu_stats", jsonData)
		if err != nil {
			log.Fatalf("failed to save cpu_stats: %v", err)
		}
	}

}
