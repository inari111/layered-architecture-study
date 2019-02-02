package api

import (
	"net/http"

	"github.com/inari111/layered-architecture-study/di"
)

func init() {
	http.Handle("/", di.InitializeAPIHandler())
}
