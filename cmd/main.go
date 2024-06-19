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

	if len(os.Args) < 4 {
		fmt.Println("need to specify directory, interval and save to database boolean (1 or 0)")
		fmt.Println("eg: go run cmd/main.go")
		os.Exit(0)
	}

	path := os.Args[1]

	if _, err := fmt.Sscanf(os.Args[2], "%d", &interval); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// if 0: don't save to database
	dbFlag := false

	if len(os.Args) > 3 && os.Args[3] == "1" {
		dbFlag = true
	}

	// more precise than sleep
	tick := time.Duration(interval) * time.Second
	ticker := time.NewTicker(tick)
	defer ticker.Stop()

	// run indefinitely
	for {
		select {
		case <-ticker.C:
			fmt.Println("tick at: ", time.Now())
			stats.MonitorDiskUsage(path, dbFlag)
			stats.MonitorProcesses(dbFlag)
			stats.MonitorCPUAndMemory(dbFlag)
		}
	}

}
