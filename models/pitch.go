package models

import (
	_ "github.com/mattn/go-sqlite3"
)

type Pitch struct {
	PitcherId        int    `json:"pitcherId"`
	BatterId         int    `json:"batterId"`
	EndTime          string `json:"endTime"`
	PitchDescription string `json:"pitchDescription"`
	CallDescription  string `json:"callDescription"`
}

func GetPitchesByPitcherId(playerId string) ([]Pitch, error) {
	rows, err := DB.Query(`
		SELECT 
			[matchup.pitcher.id],
			[matchup.batter.id],
			endTime,
			[details.type.description],
			[details.call.description]
		FROM pbp 
		WHERE [matchup.pitcher.id] = ? AND isPitch = 1
		ORDER BY endTime DESC`, playerId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	pitches := make([]Pitch, 0)

	for rows.Next() {
		pitch := Pitch{}
		err := rows.Scan(
			&pitch.PitcherId,
			&pitch.BatterId,
			&pitch.EndTime,
			&pitch.PitchDescription,
			&pitch.CallDescription,
		)

		if err != nil {
			return nil, err
		}

		pitches = append(pitches, pitch)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return pitches, nil
}

func GetPitchesByPitcherIdVsBatterId(playerId string, batterId string) ([]Pitch, error) {
	rows, err := DB.Query(`
		SELECT 
			[matchup.pitcher.id],
			[matchup.batter.id],
			endTime,
			[details.type.description],
			[details.call.description]
		FROM pbp 
		WHERE isPitch = 1
			AND [matchup.pitcher.id] = ? 
			AND [matchup.batter.id] = ?
		ORDER BY endTime DESC`, playerId, batterId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	pitches := make([]Pitch, 0)

	for rows.Next() {
		pitch := Pitch{}
		err := rows.Scan(
			&pitch.PitcherId,
			&pitch.BatterId,
			&pitch.EndTime,
			&pitch.PitchDescription,
			&pitch.CallDescription,
		)

		if err != nil {
			return nil, err
		}

		pitches = append(pitches, pitch)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return pitches, nil
}
