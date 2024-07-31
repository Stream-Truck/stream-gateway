package handler

import (
	"net/http"

	"github.com/google/wire"
)

var HandlerProviderSet = wire.NewSet(
	NewServiceList,
)

// New ServiceList
func NewServiceList() []Handler {
	return []Handler{}
}

// Service Interface
type Handler interface {
	RegisterMuxRouter(mux *http.ServeMux)
}
