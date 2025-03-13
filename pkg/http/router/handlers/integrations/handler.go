package integrations

import (
	"net/http"

	base "github.com/smarthomeix/agents/pkg/service"
	"github.com/smarthomeix/pkg/http/response"
)

type Handler struct {
	service base.ServiceInterface
}

func NewHandler(service base.ServiceInterface) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	resources := FormatResources(h.service.GetIntegrations())

	response.HandleJSON(w, resources)
}
