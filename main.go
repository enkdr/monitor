package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"time"
)

type Fsid struct {
	X__val [2]int32 `json:"fsid"`
}

type FileSystemStats struct {
	Type    int64  `json:"type"`
	Bsize   int64  `json:"block_size"`
	Blocks  uint64 `json:"total_blocks"`
	Bfree   uint64 `json:"free_blocks"`
	Bavail  uint64 `json:"available_blocks"`
	Files   uint64 `json:"total_files"`
	Ffree   uint64 `json:"free_files"`
	Fsid    Fsid
	Namelen int64    `json:"max_name_length"`
	Frsize  int64    `json:"fragment_size"`
	Flags   int64    `json:"flags"`
	Spare   [4]int64 `json:"spare"`
}

func monitorDiskUsage(path string) {

	fs := syscall.Statfs_t{}

	err := syscall.Statfs(path, &fs)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", fs)

	stats := FileSystemStats{
		Type:    fs.Type,
		Bsize:   fs.Bsize,
		Blocks:  fs.Blocks,
		Bfree:   fs.Bfree,
		Bavail:  fs.Bavail,
		Files:   fs.Files,
		Ffree:   fs.Ffree,
		Fsid:    Fsid{X__val: fs.Fsid.X__val},
		Namelen: fs.Namelen,
		Frsize:  fs.Frsize,
		Flags:   fs.Flags,
		Spare:   fs.Spare,
	}

	jsonData, err := json.MarshalIndent(stats, "", " ")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonData))
}

func monitorProcessesAndCPU() {

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
			monitorDiskUsage(path)
			monitorProcessesAndCPU()
		}
	}

	// for {
	// 	time.Sleep(tick)
	// 	fmt.Println("tick at: ", time.Now())
	// }

}
