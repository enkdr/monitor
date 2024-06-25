package main

import (
	"fmt"
	"log"
	"time"

	"github.com/enkdr/monitor/app"
	"github.com/enkdr/monitor/stats"
	_ "github.com/lib/pq"
)

func main() {

	go startApp()

	path := "/"
	interval := 2
	dbBoolean := "n"

	fmt.Println("Enter a directory to monitor (default is /)")
	fmt.Scanln(&path)
	fmt.Println("Enter an interval in seconds (default is 2)")
	fmt.Scanln(&interval)
	fmt.Println("Save to db? (y or n) default is n")
	fmt.Scanln(&dbBoolean)

	dbFlag := false

	if dbBoolean == "y" {
		dbFlag = true
	}

	// more precise than sleep
	tick := time.Duration(interval) * time.Second
	ticker := time.NewTicker(tick)
	defer ticker.Stop()

	taskChan := make(chan bool, 1)

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

func startApp() error {
	// start App
	app := app.NewApp()

	if err := app.ListenAndServe(); err != nil {
		log.Fatal("Failed to start server:", err)
		return err
	}

	log.Println("Starting server on :8080")

	return nil
}
