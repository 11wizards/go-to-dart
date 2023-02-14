package user

import (
	"time"
)

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Options   map[string]string
	Tags      []string
}
