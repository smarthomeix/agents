package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/smarthomeix/agents/pkg/http/router/handlers/integrations"
	"github.com/smarthomeix/agents/pkg/http/router/handlers/service"
	base "github.com/smarthomeix/agents/pkg/service"
)

func NewServer(port string, handlers base.ServiceInterface) *http.Server {
	mux := chi.NewMux()

	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	mountService(mux, handlers)
	mountIntegrations(mux, handlers)

	return &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}
}

func mountService(mux *chi.Mux, handlers base.ServiceInterface) {
	h := service.NewHandler(handlers)

	mux.Get("/service", h.Get)
}

func mountIntegrations(mux *chi.Mux, handlers base.ServiceInterface) {
	h := integrations.NewHandler(handlers)

	mux.Get("/integrations", h.List)
}
