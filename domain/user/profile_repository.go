package user

import "context"

type ProfileRepository interface {
	Get(ctx context.Context, id ID) (*Profile, error)
	Put(ctx context.Context, profile *Profile) error
}
