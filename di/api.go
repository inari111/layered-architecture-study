// +build wireinject

package di

import (
	"net/http"

	"github.com/inari111/layered-architecture-study/domain"
	article2 "github.com/inari111/layered-architecture-study/infra/persistence/article"

	"github.com/inari111/layered-architecture-study/application/article"
	"github.com/inari111/layered-architecture-study/handler/api"

	"github.com/google/go-cloud/wire"
)

func InitializeAPIHandler() http.Handler {
	wire.Build(
		domain.NewCurrentTimeFunc,
		article2.NewRepository,
		article.NewApplication,
		api.NewArticleService,
		api.NewHandler,
	)
	return nil
}
