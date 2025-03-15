package controller

import (
	"context"

	"github.com/smarthomeix/agents/pkg/service"
)

type ControllerDriver struct{}

func (d *ControllerDriver) GetActions() service.ActionCollection {
	return service.ActionCollection{}
}

func (d *ControllerDriver) GetTelemetry(ctx context.Context) (service.TelemetryData, error) {
	return service.TelemetryData{}, nil
}

func (d *ControllerDriver) ExecuteAction(ctx context.Context, req service.ExecuteActionRequest) (service.ExecuteActionResult, error) {
	return service.ExecuteActionResult{}, nil
}
