package service

type Configuration map[string]any

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
	NewDevice(config Configuration) (DeviceInterface, error)
}

type ParameterDefinition struct {
	Name         string
	Type         string // e.g., "string", "int", "boolean"
	Required     bool
	DefaultValue any
}

type Action struct {
	Name        string
	Description string
	Parameters  []ParameterDefinition
}

type ActionCollection map[string]Action

type Telemetry map[string]any

type ActionParameters map[string]any

type ExecuteActionRequest struct {
	Action     string
	Parameters ActionParameters
}

type ExecuteActionResult struct {
	Success bool
	Message string
}

type DeviceInterface interface {
	GetActions() ActionCollection
	GetTelemetry() (Telemetry, error)
	ExecuteAction(action ExecuteActionRequest) (ExecuteActionResult, error)
}
