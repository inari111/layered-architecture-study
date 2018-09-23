package article

import (
	"context"

	"github.com/inari111/layered-architecture-study/domain/article"
	"go.mercari.io/datastore/aedatastore"
	"go.mercari.io/datastore/boom"
)

func NewRepository() *repository {
	return &repository{}
}

type repository struct{}

func (repo *repository) Create(ctx context.Context, entity *article.Entity) error {
	client, err := aedatastore.FromContext(ctx)
	if err != nil {
		return err
	}
	b := boom.FromClient(ctx, client)
	if _, err := b.Put(fromEntity(entity)); err != nil {
		return err
	}
	return nil
}
