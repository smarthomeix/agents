package devices

import (
	base "github.com/smarthomeix/agents/pkg/service"
	"github.com/smarthomeix/pkg/validator"
)

func ValidateRequest(request CreateRequest, service base.ServiceInterface) (*base.Device, error) {
	model := &base.Device{}

	validation := validator.New()

	if err := validID(request.ID, model); err != nil {
		validation.Field("id", err)
	}

	if err := validIntegrationID(request.IntegrationID, model, service); err != nil {
		validation.Field("integrationId", err)
	}

	if len(validation) > 0 {
		return nil, validation
	}

	return model, nil
}

func validID(id string, model *base.Device) error {
	if id == "" {
		return validator.NewFieldError("ID is required")
	}

	model.ID = id

	return nil
}

func validIntegrationID(integrationID string, model *base.Device, service base.ServiceInterface) error {
	if integrationID == "" {
		return validator.NewFieldError("Integration ID is required")
	}

	integration, exists := service.GetIntegration(integrationID)

	if !exists {
		return validator.NewFieldError("Integration does not exist")
	}

	model.IntegrationID = integration.GetIntegration().ID

	return nil
}
