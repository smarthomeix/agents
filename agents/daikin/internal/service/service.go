package service

import (
	"github.com/smarthomeix/agents/agents/daikin/internal/integration/controller"
	"github.com/smarthomeix/agents/pkg/helpers/ptr"
	base "github.com/smarthomeix/agents/pkg/service"
)

type Service struct {
	service      base.Service
	integrations map[string]base.IntegrationInterface
}

func NewDaikinService() *Service {
	service := base.Service{
		ID:          "smarthomeix.service.daikin",
		Name:        "Daikin HVAC Integration",
		Version:     "1.0.0",
		Description: ptr.Str("Supports Daikin BRP072C42 series wireless controllers"),
		Brand:       ptr.Str("smarthomeIX"),
		Developer:   "smarthomeIX",
	}

	integrations := map[string]base.IntegrationInterface{
		"integration.controller": controller.NewIntegration(),
	}

	return &Service{
		service:      service,
		integrations: integrations,
	}
}

func (s *Service) GetService() base.Service {
	return s.service
}

func (s *Service) GetIntegrations() base.IntegrationCollection {
	collection := make(base.IntegrationCollection)

	for id, integration := range s.integrations {
		collection[id] = integration.GetIntegration()
	}

	return collection
}

func (s *Service) GetIntegration(integrationId string) (base.IntegrationInterface, bool) {
	integration, exists := s.integrations[integrationId]

	return integration, exists
}
