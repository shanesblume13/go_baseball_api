package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./baseball_db.db")
	if err != nil {
		return err
	}

	DB = db
	return nil
}

type Team struct {
	Id       int    `json:"team_id"`
	FullName string `json:"team_full_name"`
}

func GetTeams(count int) ([]Team, error) {
	rows, err := DB.Query("SELECT team_id, team_full_name FROM mlb_team_inf ORDER BY team_full_name ASC LIMIT ?", count)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	teams := make([]Team, 0)

	for rows.Next() {
		team := Team{}
		err := rows.Scan(&team.Id, &team.FullName)

		if err != nil {
			return nil, err
		}

		teams = append(teams, team)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return teams, nil
}
