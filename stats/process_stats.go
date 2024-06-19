package stats

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

type ProcessesData struct {
	ProcessCount int `json:"process_count"`
	CPUUsage     int `json:"cpu_usage"`
}

func MonitorProcesses(dbFlag bool) {

	psCmd := exec.Command("ps", "aux")
	output, err := psCmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
	}

	stats := ProcessesData{
		ProcessCount: bytes.Count(output, []byte("\n")),
		CPUUsage:     runtime.NumCPU(),
	}

	jsonData, err := StatsJSON(stats)

	if err != nil {
		fmt.Println(err)
	}

	if dbFlag {
		err = StatsDBInsert("process_stats", jsonData)
		if err != nil {
			log.Fatalf("failed to save process_stats: %v", err)
		}
	} else {
		fmt.Println("P R O C E S S  S T A T S")
		fmt.Println(string(jsonData))
	}

}
