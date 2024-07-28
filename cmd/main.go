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

	path := "/"
	interval := 2
	dbBoolean := "y"

	fmt.Println("Enter a directory to monitor (default is /)")
	fmt.Scanln(&path)
	fmt.Println("Enter an interval in seconds (default is 2)")
	fmt.Scanln(&interval)
	fmt.Println("Save to db? (y or n) default is y")
	fmt.Scanln(&dbBoolean)

	dbFlag := (dbBoolean == "y")

	// more precise than sleep
	tick := time.Duration(interval) * time.Second
	ticker := time.NewTicker(tick)
	defer ticker.Stop()

	taskChan := make(chan bool, 1)

	go taskWorker(path, dbFlag, taskChan)

	// Start the application server
	go func() {
		if err := startApp(); err != nil {
			log.Fatalf("Application failed to start: %v", err)
		}
	}()

	// run indefinitely
	// for {
	// 	select {
	// 	case <-ticker.C:
	// 		fmt.Println("tick at: ", time.Now())
	// 		taskChan <- true
	// 	}
	// }

	// // testing: run only 1000 times
	for x := 0; x < 1000; x++ {
		select {
		case <-ticker.C:
			fmt.Println("tick at: ", time.Now())
			fmt.Println("tick number: ", x)
			taskChan <- true
		}
	}

}

// using channels to sync DB inserts
func taskWorker(path string, dbFlag bool, taskChan <-chan bool) {

	fmt.Println("Starting taskWorkers")

	switch dbFlag {
	case true:
		fmt.Println("requesting stats and saving to the database")
	default:
		fmt.Println("requesting stats and printing to stdout only")
	}

	for range taskChan {
		stats.MonitorDiskUsage(path, dbFlag)
		stats.MonitorProcesses(dbFlag)
		stats.MonitorCPUAndMemory(dbFlag)
	}
}

func startApp() error {
	// start App
	fmt.Println("Launching application on port :8080")
	app := app.NewApp()

	if err := app.ListenAndServe(); err != nil {
		log.Fatal("Failed to start server:", err)
		return err
	}

	log.Println("Starting server on :8080")

	return nil
}
