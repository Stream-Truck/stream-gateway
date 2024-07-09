package rest_api

import (
	"application/internal/handler"
	"net/http"
)

func NewHttpHandler(
	svcs ...handler.ServiceInterface,

) http.Handler {
	mux := http.NewServeMux()

	for _, svc := range svcs {
		svc.RegisterMuxRouter(mux)
	}

	return mux

}
