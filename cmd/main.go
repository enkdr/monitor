package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/enkdr/monitor/app"
	"github.com/enkdr/monitor/stats"
	_ "github.com/lib/pq"
)

func main() {

	startApp()

	var interval int
	taskChan := make(chan bool, 1)

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

	go taskWorker(path, dbFlag, taskChan)

	// run indefinitely
	for {
		select {
		case <-ticker.C:
			fmt.Println("tick at: ", time.Now())
			taskChan <- true
		}
	}

}

// using channels to sync DB inserts
func taskWorker(path string, dbFlag bool, taskChan <-chan bool) {
	for range taskChan {
		stats.MonitorDiskUsage(path, dbFlag)
		stats.MonitorProcesses(dbFlag)
		stats.MonitorCPUAndMemory(dbFlag)
	}
}

func startApp() {
	// start App
	app := app.NewApp()

	if err := app.ListenAndServe(); err != nil {
		log.Fatal("Failed to start server:", err)
	}

	log.Println("Starting server on :8080")
}
