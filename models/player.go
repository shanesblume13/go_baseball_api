package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Player struct {
	Id                   int    `json:"person_id"`
	JerseyNumber         string `json:"jersey_number"`
	FullName             string `json:"person_full_name"`
	TeamId               int    `json:"team_id"`
	PositionCode         string `json:"position_code"`
	PositionName         string `json:"position_name"`
	PositionType         string `json:"position_type"`
	PositionAbbreviation string `json:"position_abbreviation"`
	StatusDescription    string `json:"status_description"`
	RosterType           string `json:"roster_type"`
}

func GetPlayers(count int) ([]Player, error) {
	rows, err := DB.Query(`
		SELECT 
			person_id,
			jersey_number, 
			person_full_name,
			team_id,
			position_code,
			position_name,
			position_type,
			position_abbreviation,
			status_description,
			roster_type
		FROM mlb_rosters 
		ORDER BY person_full_name ASC 
		LIMIT ?`, count)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	players := make([]Player, 0)

	for rows.Next() {
		player := Player{}
		err := rows.Scan(
			&player.Id,
			&player.JerseyNumber,
			&player.FullName,
			&player.TeamId,
			&player.PositionCode,
			&player.PositionName,
			&player.PositionType,
			&player.PositionAbbreviation,
			&player.StatusDescription,
			&player.RosterType,
		)

		if err != nil {
			return nil, err
		}

		players = append(players, player)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return players, nil
}

func GetPlayerById(id string) (Player, error) {
	stmt, err := DB.Prepare(`
		SELECT
			person_id, 
			jersey_number,
			person_full_name,
			team_id,
			position_code,
			position_name,
			position_type,
			position_abbreviation,
			status_description,
			roster_type
		FROM mlb_rosters 
		WHERE person_id = ?
		LIMIT 1`)

	if err != nil {
		return Player{}, err
	}

	player := Player{}

	sqlErr := stmt.QueryRow(id).Scan(
		&player.Id,
		&player.JerseyNumber,
		&player.FullName,
		&player.TeamId,
		&player.PositionCode,
		&player.PositionName,
		&player.PositionType,
		&player.PositionAbbreviation,
		&player.StatusDescription,
		&player.RosterType,
	)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return Player{}, nil
		}
		return Player{}, sqlErr
	}

	return player, nil
}

func GetPlayersByTeamId(teamId string) ([]Player, error) {
	rows, err := DB.Query(`
		SELECT 
			person_id,
			jersey_number, 
			person_full_name,
			team_id,
			position_code,
			position_name,
			position_type,
			position_abbreviation,
			status_description,
			roster_type
		FROM mlb_rosters 
		WHERE team_id = ?
		ORDER BY person_full_name ASC `, teamId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	players := make([]Player, 0)

	for rows.Next() {
		player := Player{}
		err := rows.Scan(
			&player.Id,
			&player.JerseyNumber,
			&player.FullName,
			&player.TeamId,
			&player.PositionCode,
			&player.PositionName,
			&player.PositionType,
			&player.PositionAbbreviation,
			&player.StatusDescription,
			&player.RosterType,
		)

		if err != nil {
			return nil, err
		}

		players = append(players, player)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return players, nil
}
