package user

import (
	"time"
)

type User struct {
	Profile
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Options   map[string]string
	Tags      []string
}

type Profile struct {
	ID    int
	Name  string `json:"full_name"`
	Email string
	Address
}

type Address struct {
	Street string `json:"street_line_1"`
	City   string
	State  string
}
