package devices

import (
	"time"

	base "github.com/smarthomeix/agents/pkg/service"
)

type Resource struct {
	ID            string         `json:"id"`
	IntegrationID string         `json:"integrationID"`
	Config        map[string]any `json:"config"`
	RegisteredAt  string         `json:"registeredAt"`
}

func FormatResource(model base.Device) Resource {
	return Resource{
		ID:            model.ID,
		IntegrationID: model.IntegrationID,
		Config:        model.Config,
		RegisteredAt:  model.RegisteredAt.UTC().Format(time.RFC3339),
	}
}
