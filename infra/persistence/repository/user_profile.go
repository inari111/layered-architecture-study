package repository

import (
	"context"

	"github.com/inari111/layered-architecture-study/domain/user"
	"github.com/inari111/layered-architecture-study/infra/persistence/datastore"
	"go.mercari.io/datastore/aedatastore"
	"go.mercari.io/datastore/boom"
)

func NewUserProfileRepository() user.ProfileRepository {
	return &userProfileRepository{}
}

type userProfileRepository struct{}

func (*userProfileRepository) Get(ctx context.Context, id user.ID) (*user.Profile, error) {
	client, err := aedatastore.FromContext(ctx)
	if err != nil {
		return nil, err
	}
	b := boom.FromClient(ctx, client)
	u := datastore.User{ID: id.String()}
	if err := b.Get(&u); err != nil {
		return nil, err
	}
	return u.ToDomainProfile(), nil
}

func (*userProfileRepository) Put(ctx context.Context, profile *user.Profile) error {
	client, err := aedatastore.FromContext(ctx)
	if err != nil {
		return err
	}
	b := boom.FromClient(ctx, client)
	if _, err := b.Put(datastore.NewUser().WithDomainProfile(profile)); err != nil {
		return err
	}
	return nil
}
