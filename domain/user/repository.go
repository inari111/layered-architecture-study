package user

import "context"

type Repository interface {
	Get(ctx context.Context, id ID) (*User, error)
	Put(ctx context.Context, u *User) error
}
