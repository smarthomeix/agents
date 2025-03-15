package devices

import (
	"time"

	"github.com/smarthomeix/agents/pkg/director"
	"github.com/smarthomeix/agents/pkg/service"
)

type Resource struct {
	ID            string               `json:"id"`
	IntegrationID string               `json:"integrationID"`
	Config        service.DeviceConfig `json:"config"`
	Telemetry     TelemetryResource    `json:"telemetry"`
	RegisteredAt  string               `json:"registeredAt"`
}

type TelemetryResource struct {
	Data      service.TelemetryData `json:"data"`
	UpdatedAt *string               `json:"updatedAt"`
}

func FormatResource(model *director.Device) Resource {
	telemetry := TelemetryResource{
		Data: model.Telemetry.Data,
	}

	if model.Telemetry.UpdatedAt != nil {
		updatedAt := model.Telemetry.UpdatedAt.UTC().Format(time.RFC3339)

		telemetry.UpdatedAt = &updatedAt
	}

	return Resource{
		ID:            model.ID,
		IntegrationID: model.IntegrationID,
		Config:        model.Config,
		RegisteredAt:  model.RegisteredAt.UTC().Format(time.RFC3339),
		Telemetry:     telemetry,
	}
}
