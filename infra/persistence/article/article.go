package article

import (
	"time"

	"github.com/inari111/layered-architecture-study/doamin/article"
)

type Article struct {
	ID        string `datastore:"-" boom:"id"`
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func toEntity(v *Article) *article.Entity {
	return &article.Entity{
		ID:        article.ID(v.ID),
		Title:     v.Title,
		Body:      v.Body,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}
}

func fromEntity(v *article.Entity) *Article {
	return &Article{
		Title:     v.Title,
		Body:      v.Body,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}
}
