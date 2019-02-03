package api

import (
	"net/http"

	"github.com/inari111/layered-architecture-study/handler/rpc"

	"github.com/inari111/layered-architecture-study/handler"
	"github.com/twitchtv/twirp"
)

type server struct {
	pathPrefix string
	twServer   pb.TwirpServer
}

func NewHandler(
	articleService pb.ArticleService,
	userService pb.UserService,
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
		{
			pathPrefix: pb.UserServicePathPrefix,
			twServer: pb.NewUserServiceServer(
				userService,
				hooks,
			),
		},
	}

	for _, s := range servers {
		mux.Handle(s.pathPrefix, handler.InjectAppEngineContext(s.twServer))
	}
	return mux
}
