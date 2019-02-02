package article

import (
	"time"

	"github.com/inari111/layered-architecture-study/domain/article"
)

type Article struct {
	ID        string `datastore:"-" boom:"id"`
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (a *Article) toDomain() *article.Article {
	return &article.Article{
		ID:        article.ID(a.ID),
		Title:     a.Title,
		Body:      a.Body,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}

func articleFrom(a *article.Article) *Article {
	return &Article{
		Title:     a.Title,
		Body:      a.Body,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}
