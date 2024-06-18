package stats

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
)

func MonitorProcesses() {

	psCmd := exec.Command("ps", "aux")
	output, err := psCmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
	}

	processCount := bytes.Count(output, []byte("\n"))

	cpuUsage := runtime.NumCPU()

	fmt.Printf("Number of processes: %d\n", processCount)
	fmt.Printf("CPU usage: %d\n", cpuUsage)

}
