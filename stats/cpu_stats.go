package stats

import (
	"fmt"
	"runtime"
)

func MonitorCPU() {

	cpuUsage := runtime.NumCPU()

	fmt.Printf("CPU usage: %d\n", cpuUsage)

}
