package devices

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
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

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {

	request := CreateRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		response.HandleStatus(w, http.StatusBadRequest)
		return
	}

	service := h.director.GetService()

	model, err := ValidateRequest(request, service)

	if err != nil {
		response.HandleValidationError(w, err)
		return
	}

	err = h.director.Attach(model)

	if err != nil {
		response.HandleStatus(w, http.StatusConflict)
		return
	}

	resources := FormatResource(*model)

	response.HandleJSON(w, resources)

}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "deviceId")

	h.director.Detach(id)

	response.HandleStatus(w, http.StatusNoContent)
}
