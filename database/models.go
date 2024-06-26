package database

import (
	"encoding/json"
	"time"
)

type StatsData struct {
	ID        int             `json:"id" db:"id"`
	StatsJSON json.RawMessage `json:"stats_json" db:"stats_json"`
	CreatedAt time.Time       `json:"created_at" db:"created_at"`
}
