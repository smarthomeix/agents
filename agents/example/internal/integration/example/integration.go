package example

import (
	"github.com/smarthomeix/agents/pkg/helpers/ptr"
	"github.com/smarthomeix/agents/pkg/service"
)

type ControllerIntegration struct{}

func NewIntegration() *ControllerIntegration {
	return &ControllerIntegration{}
}

func (i *ControllerIntegration) GetIntegration() service.Integration {
	return service.Integration{
		ID:          "integration.example",
		Name:        "Example integration",
		Description: ptr.Str("Integration for testing and development purposes"),
	}
}

func (i *ControllerIntegration) NewDriver(config service.DeviceConfig) (service.DriverInterface, error) {
	return &ExampleDriver{}, nil
}
