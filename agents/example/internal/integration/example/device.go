package example

import (
	"context"

	"github.com/smarthomeix/agents/pkg/service"
)

type ExampleDriver struct{}

func (d *ExampleDriver) GetActions() service.ActionCollection {
	return service.ActionCollection{}
}

func (d *ExampleDriver) GetTelemetry(ctx context.Context) (service.TelemetryData, error) {
	return service.TelemetryData{
		"mode":        "cool",
		"temperature": 25,
	}, nil
}

func (d *ExampleDriver) ExecuteAction(ctx context.Context, req service.ExecuteActionRequest) (service.ExecuteActionResult, error) {
	return service.ExecuteActionResult{}, nil
}
