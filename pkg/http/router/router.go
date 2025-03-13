package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/smarthomeix/agents/pkg/director"
	"github.com/smarthomeix/agents/pkg/http/router/handlers/devices"
	"github.com/smarthomeix/agents/pkg/http/router/handlers/integrations"
	"github.com/smarthomeix/agents/pkg/http/router/handlers/service"
)

func NewServer(port string, director *director.Director) *http.Server {
	mux := chi.NewMux()

	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	mountService(mux, director)
	mountIntegrations(mux, director)
	mountDevices(mux, director)

	return &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}
}

func mountService(mux *chi.Mux, director *director.Director) {
	h := service.NewHandler(director)

	mux.Get("/service", h.Get)
}

func mountIntegrations(mux *chi.Mux, director *director.Director) {
	h := integrations.NewHandler(director)

	mux.Get("/integrations", h.List)
}

func mountDevices(mux *chi.Mux, director *director.Director) {
	h := devices.NewHandler(director)

	mux.Post("/devices", h.Post)

	mux.With(h.Middleware).Delete("/devices/{deviceId}", h.Delete)
}
