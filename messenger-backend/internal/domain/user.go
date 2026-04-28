package domain

import "time"

type User struct {
	ID           uuid.UUID
	Login        string
	PasswordHash string
	Nickname     string
	CreatedAt    time.Time
}
