package datastore

import (
	"time"

	"github.com/inari111/layered-architecture-study/domain/user"
)

type User struct {
	ID string `datastore:"-" boom:"id"`

	Name        string
	Age         int
	Description string `datastore:",noindex"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) ToDomainUser() *user.User {
	return &user.User{
		ID:        user.ID(u.ID),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (u *User) ToDomainProfile() *user.Profile {
	return &user.Profile{
		Name:        u.Name,
		Age:         u.Age,
		Description: u.Description,
	}
}

func (u *User) WithDomainUser(du *user.User) *User {
	dsUser := *u
	dsUser.ID = du.ID.String()
	dsUser.CreatedAt = du.CreatedAt
	dsUser.UpdatedAt = du.UpdatedAt
	return &dsUser
}

func (u *User) WithDomainProfile(dp *user.Profile) *User {
	dsUser := *u
	dsUser.Name = dp.Name
	dsUser.Age = dp.Age
	dsUser.Description = dp.Description
	return &dsUser
}

func NewUser() *User {
	return &User{}
}
