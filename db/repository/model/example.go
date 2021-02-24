package model

import "time"

type Example struct {
	ID uint
	CreatedAt time.Time
	DeletedAt *time.Time
	UpdatedAt time.Time
}
