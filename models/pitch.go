package models

import (
	_ "github.com/mattn/go-sqlite3"
)

type Pitch struct {
	PitcherId       int    `json:"matchup.pitcher.id"`
	EndTime         string `json:"endTime"`
	CallDescription string `json:"details.call.description"`
}

func GetPitchesByPlayerId(playerId string) ([]Pitch, error) {
	rows, err := DB.Query(`
		SELECT 
			[matchup.pitcher.id],
			endTime,
			[details.call.description]
		FROM pbp 
		WHERE [matchup.pitcher.id] = ? AND isPitch = 1
		ORDER BY endTime DESC
		LIMIT 100`, playerId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	pitches := make([]Pitch, 0)

	for rows.Next() {
		pitch := Pitch{}
		err := rows.Scan(
			&pitch.PitcherId,
			&pitch.EndTime,
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
