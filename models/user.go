package models

import "time"

type User struct {
	ID        uint
	Name      string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
