package article

import "time"

type ID string

func (id ID) String() string {
	return string(id)
}

type Article struct {
	ID        ID
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
