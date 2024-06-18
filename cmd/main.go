package main

import (
	"fmt"
	"os"
	"time"

	"github.com/enkdr/monitor/stats"
	_ "github.com/lib/pq"
)

func main() {
	var interval int

	if len(os.Args) < 3 {
		fmt.Println("need to specify directory and interval")
		os.Exit(0)
	}

	path := os.Args[1]

	if _, err := fmt.Sscanf(os.Args[2], "%d", &interval); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	tick := time.Duration(interval) * time.Second

	// more precise than sleep
	ticker := time.NewTicker(tick)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("tick at: ", time.Now())
			stats.MonitorDiskUsage(path)
			stats.MonitorProcessesAndCPU()
		}
	}

}
