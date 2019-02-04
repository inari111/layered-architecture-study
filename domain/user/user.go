package user

import "time"

type ID string

func (id ID) String() string {
	return string(id)
}

type User struct {
	ID ID

	Profile

	CreatedAt time.Time
	UpdatedAt time.Time
}
