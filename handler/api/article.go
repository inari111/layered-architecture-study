package api

import (
	"context"

	articleApp "github.com/inari111/layered-architecture-study/application/article"

	"github.com/inari111/layered-architecture-study/handler/rpc"
)

func NewArticleService(app articleApp.Application) pb.ArticleService {
	return &articleService{
		app: app,
	}
}

type articleService struct {
	app articleApp.Application
}

func (s *articleService) Create(ctx context.Context, req *pb.ArticleServiceCreateRequest) (*pb.ArticleServiceCreateResponsee, error) {
	entity, err := s.app.Create(ctx, req.GetTitle(), req.GetBody())
	if err != nil {
		return nil, err
	}
	return &pb.ArticleServiceCreateResponsee{
		Article: &pb.Article{
			Id:    entity.ID.String(),
			Title: entity.Title,
			Body:  entity.Body,
		},
	}, nil
}
