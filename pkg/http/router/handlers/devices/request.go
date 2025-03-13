package devices

type CreateRequest struct {
	ID            string         `json:"id"`
	IntegrationID string         `json:"integrationId"`
	Config        map[string]any `json:"config"`
}
