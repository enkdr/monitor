package database

import "time"

type Plant struct {
	ID          int       `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	Image       string    `db:"image" json:"image"`
	Water       int       `db:"water" json:"water"`             // milliliters per day
	Sun         int       `db:"sun" json:"sun"`                 // hours of sunshine per day
	Germination int       `db:"germination" json:"germination"` // days to germinate
	Flowering   int       `db:"flowering" json:"flowering"`     // days to flower
	Harvest     int       `db:"harvest" json:"harvest"`         // days to harvest
	Seed        int       `db:"seed" json:"seed"`               // seeds per plant
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
