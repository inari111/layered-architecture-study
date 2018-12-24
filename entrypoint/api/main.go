package api

import (
	"net/http"

	"github.com/inari111/layered-architecture-study/di"
	"github.com/inari111/layered-architecture-study/handler/rpc"
)

type server struct {
	pathPrefix string
	twServer   pb.TwirpServer
}

func init() {
	http.Handle("/", di.InitializeAPIHandler())
}
