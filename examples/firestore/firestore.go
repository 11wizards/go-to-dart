package firestore

import (
	"time"
)

type User struct {
	ID        int               `firestore:"id"`
	Name      string            `firestore:"name"`
	Email     string            `firestore:"email"`
	Password  string            `firestore:"password"`
	CreatedAt time.Time         `firestore:"createdAt"`
	UpdatedAt time.Time         `firestore:"updatedAt"`
	DeletedAt *time.Time        `firestore:"deletedAt,omitempty"`
	Options   map[string]string `firestore:"options,omitempty"`
	Tags      []string          `firestore:"tags,omitempty"`
}
