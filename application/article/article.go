package article

import (
	"context"

	"github.com/inari111/layered-architecture-study/domain"
	"github.com/inari111/layered-architecture-study/domain/article"
)

type Application interface {
	Create(ctx context.Context, title, body string) (*article.Article, error)
}

func NewApplication(repo article.Repository, now domain.CurrentTimeFunc) Application {
	return &articleAppImpl{
		repo: repo,
		now:  now,
	}
}

type articleAppImpl struct {
	repo article.Repository
	now  domain.CurrentTimeFunc
}

func (app *articleAppImpl) Create(ctx context.Context, title, body string) (*article.Article, error) {
	a := article.Article{
		Title:     title,
		Body:      body,
		CreatedAt: app.now(),
		UpdatedAt: app.now(),
	}
	if err := app.repo.Create(ctx, &a); err != nil {
		return nil, err
	}
	return &a, nil
}
