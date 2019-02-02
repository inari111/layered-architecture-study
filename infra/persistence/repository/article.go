package repository

import (
	"context"

	"github.com/inari111/layered-architecture-study/infra/persistence/datastore"

	"github.com/inari111/layered-architecture-study/domain/article"
	"go.mercari.io/datastore/aedatastore"
	"go.mercari.io/datastore/boom"
)

func NewArticleRepository() article.Repository {
	return &articleRepository{}
}

type articleRepository struct{}

func (repo *articleRepository) Create(ctx context.Context, article *article.Article) error {
	client, err := aedatastore.FromContext(ctx)
	if err != nil {
		return err
	}
	b := boom.FromClient(ctx, client)
	if _, err := b.Put(datastore.ArticleFrom(article)); err != nil {
		return err
	}
	return nil
}
