package stats

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/enkdr/monitor/config"
	"github.com/enkdr/monitor/database"
)

func StatsJSON(stats interface{}) ([]byte, error) {
	jsonData, err := json.MarshalIndent(stats, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return nil, err
	}

	fmt.Println(string(jsonData))

	return jsonData, nil
}

func StatsDBInsert(statsJson []byte) error {

	config, err := config.GetConfig()
	if err != nil {
		return err
	}

	db, err := database.InitDB(config)
	if err != nil {
		return err
	}

	defer db.Close()

	qry := `INSERT INTO public.fs_stats (fs_stats_json, created_at) VALUES($1, $2);`
	_, err = db.Exec(qry, statsJson, time.Now())
	if err != nil {
		return err
	}

	return nil

}
