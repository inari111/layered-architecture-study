// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"github.com/inari111/layered-architecture-study/application/article"
	"github.com/inari111/layered-architecture-study/domain"
	"github.com/inari111/layered-architecture-study/handler/api"
	"github.com/inari111/layered-architecture-study/infra/persistence/repository"
	"net/http"
)

// Injectors from api.go:

func InitializeAPIHandler() http.Handler {
	articleRepository := repository.NewArticleRepository()
	currentTimeFunc := domain.NewCurrentTimeFunc()
	application := article.NewApplication(articleRepository, currentTimeFunc)
	articleService := api.NewArticleService(application)
	userService := api.NewUserService()
	handler := api.NewHandler(articleService, userService)
	return handler
}
