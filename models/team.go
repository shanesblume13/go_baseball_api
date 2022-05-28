package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Team struct {
	Id            int    `json:"id"`
	Abbreviation  string `json:"abbreviation"`
	FullName      string `json:"fullName"`
	ShortName     string `json:"shortName"`
	FranchiseName string `json:"franchiseName"`
	ClubName      string `json:"clubName"`
	VenueId       int    `json:"venueId"`
	VenueName     string `json:"venueName"`
	LeagueId      int    `json:"leagueId"`
	LeagueName    string `json:"leagueName"`
	DivisionId    int    `json:"divisionId"`
	DivisionName  string `json:"divisionName"`
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
