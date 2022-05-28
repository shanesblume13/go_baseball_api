package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Team struct {
	Id            int    `json:"team_id"`
	Abbreviation  string `json:"team_abbreviation"`
	FullName      string `json:"team_full_name"`
	ShortName     string `json:"short_name"`
	FranchiseName string `json:"franchise_name"`
	ClubName      string `json:"club_name"`
	VenueId       int    `json:"venue_id"`
	VenueName     string `json:"venue_name"`
	LeagueId      int    `json:"league_id"`
	LeagueName    string `json:"league_name"`
	DivisionId    int    `json:"division_id"`
	DivisionName  string `json:"division_name"`
}

func GetTeams(count int) ([]Team, error) {
	rows, err := DB.Query(`
		SELECT 
			team_id,
			team_abbreviation, 
			team_full_name,
			short_name,
			franchise_name,
			club_name,
			venue_id,
			venue_name,
			league_id,
			league_name,
			division_id,
			division_name
		FROM mlb_team_inf 
		ORDER BY team_full_name ASC 
		LIMIT ?`, count)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	teams := make([]Team, 0)

	for rows.Next() {
		team := Team{}
		err := rows.Scan(&team.Id, &team.Abbreviation, &team.FullName, &team.ShortName, &team.FranchiseName, &team.ClubName, &team.VenueId, &team.VenueName, &team.LeagueId, &team.LeagueName, &team.DivisionId, &team.DivisionName)

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

func GetTeamById(id string) (Team, error) {
	stmt, err := DB.Prepare(`
		SELECT
			team_id, 
			team_abbreviation,
			team_full_name,
			short_name,
			franchise_name,
			club_name,
			venue_id,
			venue_name,
			league_id,
			league_name,
			division_id,
			division_name
		FROM mlb_team_inf 
		WHERE team_id = ?
		LIMIT 1`)

	if err != nil {
		return Team{}, err
	}

	team := Team{}

	sqlErr := stmt.QueryRow(id).Scan(&team.Id, &team.Abbreviation, &team.FullName, &team.ShortName, &team.FranchiseName, &team.ClubName, &team.VenueId, &team.VenueName, &team.LeagueId, &team.LeagueName, &team.DivisionId, &team.DivisionName)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return Team{}, nil
		}
		return Team{}, sqlErr
	}

	return team, nil

}
