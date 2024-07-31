package model

import "time"

type Base struct {
	ID        int32
	CreatedAt time.Time
	UpdatedAt time.Time
}
