package api

import (
	"context"
	"net/http"

	articleApp "github.com/inari111/layered-architecture-study/application/article"
	"github.com/twitchtv/twirp"

	"github.com/inari111/layered-architecture-study/handler"
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

type server struct {
	pathPrefix string
	twServer   pb.TwirpServer
}

func NewHandler(
	articleService pb.ArticleService,
) http.Handler {
	mux := http.NewServeMux()

	hooks := twirp.ChainHooks(
		handler.UserAuthHooks(),
		handler.LoggingErrorHooks(),
	)

	servers := []server{
		{
			pathPrefix: pb.ArticleServicePathPrefix,
			twServer: pb.NewArticleServiceServer(
				articleService,
				hooks,
			),
		},
	}

	for _, s := range servers {
		mux.Handle(s.pathPrefix, handler.InjectAppEngineContext(s.twServer))
	}
	return mux
}
