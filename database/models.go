package database

import (
	"time"
)

type Data struct {
	ID        int       `json:"id" db:"id"`
	StatsJSON []byte    `json:"stats_json" db:"stats_json"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
