// +build wireinject

package di

import (
	"net/http"

	"github.com/inari111/layered-architecture-study/domain"
	"github.com/inari111/layered-architecture-study/infra/persistence/repository"

	"github.com/inari111/layered-architecture-study/application/article"
	"github.com/inari111/layered-architecture-study/handler/api"

	"github.com/google/go-cloud/wire"
)

func InitializeAPIHandler() http.Handler {
	wire.Build(
		domain.NewCurrentTimeFunc,
		repository.NewArticleRepository,
		article.NewApplication,
		api.NewArticleService,
		api.NewHandler,
	)
	return nil
}
