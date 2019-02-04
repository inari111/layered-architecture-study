package repository

import (
	"context"

	"github.com/inari111/layered-architecture-study/infra/persistence/datastore"

	"github.com/inari111/layered-architecture-study/domain/user"
	"go.mercari.io/datastore/aedatastore"
	"go.mercari.io/datastore/boom"
)

func NewUserRepository() user.Repository {
	return &userRepository{}
}

type userRepository struct{}

func (*userRepository) Get(ctx context.Context, id user.ID) (*user.User, error) {
	client, err := aedatastore.FromContext(ctx)
	if err != nil {
		return nil, err
	}
	b := boom.FromClient(ctx, client)
	u := datastore.User{ID: id.String()}
	if err := b.Get(&u); err != nil {
		return nil, err
	}
	return u.ToDomainUser(), nil
}

func (*userRepository) Put(ctx context.Context, u *user.User) error {
	client, err := aedatastore.FromContext(ctx)
	if err != nil {
		return err
	}
	b := boom.FromClient(ctx, client)
	if _, err := b.Put(datastore.NewUser().WithDomainUser(u)); err != nil {
		return err
	}
	return nil
}
