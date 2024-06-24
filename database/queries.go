package database

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// move this from stats/utils -- have to inject db
// func InsertStatsData(tableName string, statsJson []byte) error {

// 	qry := fmt.Sprintf(`INSERT INTO public.%s (stats_json, created_at) VALUES($1, $2);`, tableName)
// 	_, err = db.Exec(qry, statsJson, time.Now())
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println("Inserting: %s", tableName)

// 	return nil

// }

func GetStatsData(db *sqlx.DB, table string) ([]Data, error) {
	query := fmt.Sprintf("SELECT id, stats_json, created_at FROM %s", table)

	rows, err := db.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []Data
	for rows.Next() {
		var d Data
		if err := rows.StructScan(&d); err != nil {
			return nil, err
		}
		data = append(data, d)
	}

	// Check for errors after the loop completes
	if err := rows.Err(); err != nil {
		return nil, err
	}

	fmt.Println(data)

	return data, nil
}

func GetRecentStatsData(db *sqlx.DB, table string) (*Data, error) {

	var data Data
	query := fmt.Sprintf("SELECT id, stats_json, created_at FROM %s ORDER BY created_at DESC LIMIT 1", table)
	err := db.Get(&data, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no records found in table %s", table)
		}
		return nil, err
	}
	return &data, nil
}
