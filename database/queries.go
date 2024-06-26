package database

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func GetStatsData(db *sqlx.DB, table string) ([]StatsData, error) {
	query := fmt.Sprintf("SELECT id, stats_json, created_at FROM %s", table)

	rows, err := db.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []StatsData
	for rows.Next() {
		var d StatsData
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

func GetRecentStatsData(db *sqlx.DB, table string) (*StatsData, error) {

	var data StatsData
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
