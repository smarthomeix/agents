package service

import (
	"context"
	"time"
)

type Service struct {
	ID          string
	Name        string
	Version     string
	Description *string
	Brand       *string
	Developer   string
}

type ServiceInterface interface {
	GetService() Service
	GetIntegrations() IntegrationCollection
	GetIntegration(integrationId string) (IntegrationInterface, bool)
}

type Integration struct {
	ID          string
	Name        string
	Description *string
}

type IntegrationCollection map[string]Integration

type IntegrationInterface interface {
	GetIntegration() Integration
	NewDriver(config DeviceConfig) (DriverInterface, error)
}

type Action struct {
	Name        string
	Description string
}

type ActionCollection map[string]Action

type TelemetryData map[string]any

type Telemetry struct {
	Data      TelemetryData
	UpdatedAt *time.Time
}

type ActionParameters map[string]any

type ExecuteActionRequest struct {
	Action     string
	Parameters ActionParameters
}

type ExecuteActionResult struct {
	Success bool
	Message string
}

type DriverInterface interface {
	GetActions() ActionCollection
	GetTelemetry(ctx context.Context) (TelemetryData, error)
	ExecuteAction(ctx context.Context, action ExecuteActionRequest) (ExecuteActionResult, error)
}

type DeviceConfig map[string]any
