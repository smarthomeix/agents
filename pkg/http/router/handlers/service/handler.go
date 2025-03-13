package service

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

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	resource := FormatResource(h.service.GetService())

	response.HandleJSON(w, resource)
}
