package messages

import "github.com/smarthomeix/agents/pkg/service"

type Telemetry struct {
	DeviceID      string                `json:"deviceId"`
	Source        string                `json:"source"`
	IntegrationID string                `json:"integrationId"`
	Telemetry     service.TelemetryData `json:"telemetry"`
	PublishedAt   string                `json:"publishedAt"`
	UpdatedAt     *string               `json:"updatedAt"`
}
