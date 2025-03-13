package integrations

import (
	"net/http"

	"github.com/smarthomeix/agents/pkg/director"
	"github.com/smarthomeix/pkg/http/response"
)

type Handler struct {
	director *director.Director
}

func NewHandler(director *director.Director) *Handler {
	return &Handler{
		director: director,
	}
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	service := h.director.GetService()

	resources := FormatResources(service.GetIntegrations())

	response.HandleJSON(w, resources)
}
