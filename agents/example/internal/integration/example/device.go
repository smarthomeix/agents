package example

import (
	"github.com/smarthomeix/agents/pkg/service"
)

type ExampleDriver struct{}

func (d *ExampleDriver) GetActions() service.ActionCollection {
	return service.ActionCollection{}
}

func (d *ExampleDriver) GetTelemetry() (service.Telemetry, error) {
	return service.Telemetry{}, nil
}

func (d *ExampleDriver) ExecuteAction(req service.ExecuteActionRequest) (service.ExecuteActionResult, error) {
	return service.ExecuteActionResult{}, nil
}
