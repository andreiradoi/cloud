package data

import (
	"time"
)

type Input struct {
	ID           int32  `db:"id,omitempty"`
	ExpeditionID int32  `db:"expedition_id"`
	Name         string `db:"name"`
	TeamID       *int32 `db:"team_id"`
	UserID       *int32 `db:"user_id"`
	Active       bool   `db:"active"`
}

type InputToken struct {
	ID           int32     `db:"id,omitempty"`
	ExpeditionID int32     `db:"expedition_id"`
	Token        Token     `db:"token"`
	Expires      time.Time `db:"time"`
}
