package database

import (
	"encoding/json"
	"log"

	"github.com/jmoiron/sqlx"
)

// called from server routes
func PlantQuery(db *sqlx.DB) []byte {


    var plants []Plant
    err := db.Select(&plants, "SELECT id, name, description, image, water, sun, germination, flowering, harvest, seed, created_at, updated_at from plant") // Adjust the query according to your table name and schema
    if err != nil {
        log.Fatalln(err)
    }

	
	plantsJSON, err := json.Marshal(plants)

	if err != nil {
		return []byte("error connecting to the database")
	}

	return plantsJSON

}
