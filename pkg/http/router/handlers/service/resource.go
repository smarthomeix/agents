package service

import base "github.com/smarthomeix/agents/pkg/service"

type Resource struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Version     string  `json:"version"`
	Description *string `json:"description"`
	Brand       *string `json:"brand"`
	Developer   string  `json:"developer"`
}

func FormatResource(model base.Service) Resource {
	return Resource{
		ID:          model.ID,
		Name:        model.Name,
		Version:     model.Version,
		Description: model.Description,
		Brand:       model.Brand,
		Developer:   model.Developer,
	}
}
