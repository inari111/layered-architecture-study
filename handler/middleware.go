package handler

import (
	"context"
	"net/http"

	"github.com/twitchtv/twirp"
	"google.golang.org/appengine"
)

const requestKey = "__request__"

func InjectAppEngineContext(base http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.WithContext(r.Context(), r)
		context.WithValue(ctx, requestKey, r)
		base.ServeHTTP(w, r.WithContext(ctx))
	})
}

func UserAuthHooks() *twirp.ServerHooks {
	return &twirp.ServerHooks{
		// TODO
		RequestReceived: func(ctx context.Context) (context.Context, error) {
			return ctx, nil
		},
	}
}

func LoggingErrorHooks() *twirp.ServerHooks {
	return &twirp.ServerHooks{
		// TODO
		Error: func(ctx context.Context, twerr twirp.Error) context.Context {
			return ctx
		},
	}
}
