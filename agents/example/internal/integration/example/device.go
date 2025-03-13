package example

import (
	"github.com/smarthomeix/agents/pkg/service"
)

type ControllerDevice struct{}

func (d *ControllerDevice) GetActions() service.ActionCollection {
	return service.ActionCollection{}
}

func (d *ControllerDevice) GetTelemetry() (service.Telemetry, error) {
	return service.Telemetry{}, nil
}

func (d *ControllerDevice) ExecuteAction(req service.ExecuteActionRequest) (service.ExecuteActionResult, error) {
	return service.ExecuteActionResult{}, nil
}
