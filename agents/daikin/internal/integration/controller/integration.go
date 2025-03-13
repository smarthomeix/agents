package controller

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
		ID:          "integration.controller",
		Name:        "BRP072C42 Wireless Controller",
		Description: ptr.Str("Integration for Daikin Air Conditioners"),
	}
}

func (i *ControllerIntegration) NewDevice(config service.Configuration) (service.DriverInterface, error) {
	return &ControllerDriver{}, nil
}
