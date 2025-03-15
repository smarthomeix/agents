package devices

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/smarthomeix/agents/pkg/director"
	"github.com/smarthomeix/pkg/http/response"
)

type contextKey string

const ContextKey contextKey = "device"

func (h *Handler) Middleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		deviceID := chi.URLParam(r, "deviceId")

		ctx := r.Context()

		record, exists := h.director.GetDevice(deviceID)

		if !exists {
			response.HandleNotFound(w)
			return
		}

		ctxWithDevice := context.WithValue(ctx, ContextKey, record)

		next.ServeHTTP(w, r.WithContext(ctxWithDevice))
	})
}

func GetDeviceFromContext(r *http.Request) (*director.Device, bool) {
	device, ok := r.Context().Value(ContextKey).(*director.Device)

	return device, ok
}
