package api

import (
	"net/http"
	"time"

	"github.com/inari111/layered-architecture-study/infra/persistence/article"

	articleApp "github.com/inari111/layered-architecture-study/application/article"
	"github.com/inari111/layered-architecture-study/handler/api"
	"github.com/inari111/layered-architecture-study/handler/rpc"
	"github.com/twitchtv/twirp"
)

type server struct {
	pathPrefix string
	twServer   pb.TwirpServer
}

func init() {
	mux := http.DefaultServeMux

	hooks := twirp.ChainHooks(
		userAuthHooks(),
		loggingErrorHooks(),
	)

	servers := []server{
		{
			pathPrefix: pb.ArticleServicePathPrefix,
			twServer: pb.NewArticleServiceServer(
				api.NewArticleService(
					articleApp.NewApplication(
						article.NewRepository(),
						time.Now(), // TODO
					),
				),
				hooks,
			),
		},
	}

	for _, s := range servers {
		mux.Handle(s.pathPrefix, injectAppEngineContext(s.twServer))
	}

}
