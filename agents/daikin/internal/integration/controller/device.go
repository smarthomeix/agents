package controller

import (
	"github.com/smarthomeix/agents/pkg/service"
)

type ControllerDriver struct{}

func (d *ControllerDriver) GetActions() service.ActionCollection {
	return service.ActionCollection{}
}

func (d *ControllerDriver) GetTelemetry() (service.Telemetry, error) {
	return service.Telemetry{}, nil
}

func (d *ControllerDriver) ExecuteAction(req service.ExecuteActionRequest) (service.ExecuteActionResult, error) {
	return service.ExecuteActionResult{}, nil
}
