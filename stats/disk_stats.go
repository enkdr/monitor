package stats

import (
	"encoding/json"
	"fmt"
	"log"
	"syscall"
	"time"

	"github.com/enkdr/monitor/config"
	"github.com/enkdr/monitor/database"
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

type FsData struct {
	FileSystemStats `json:"fs_stats"`
	CreatedAt       time.Time `json:"created_at"`
}

func MonitorDiskUsage(path string) {

	fs := syscall.Statfs_t{}

	err := syscall.Statfs(path, &fs)
	if err != nil {
		fmt.Println(err)
	}

	stats := FsData{
		FileSystemStats: FileSystemStats{
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
		},
	}

	jsonFsData, err := json.MarshalIndent(stats, "", " ")

	if err != nil {
		fmt.Println(err)
	}

	config, err := config.GetConfig()
	if err != nil {
		log.Fatalln("Failed to retrieve configs:", err)
	}

	db, err := database.InitDB(config)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	defer db.Close()

	qry := `INSERT INTO public.fs_stats (fs_stats_json, created_at) VALUES($1, $2);`
	_, err = db.Exec(qry, jsonFsData, time.Now())
	if err != nil {
		log.Fatalf("failed to insert data: %v", err)
	}

	// save to database
	fmt.Println(string(jsonFsData))
}
