package integrations

import base "github.com/smarthomeix/agents/pkg/service"

type Resource struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type ResourceCollection []Resource

type ResourceList struct {
	Integrations ResourceCollection `json:"integrations"`
}

func FormatResource(model base.Integration) Resource {
	return Resource{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
	}
}

func FormatResources(collection base.IntegrationCollection) ResourceList {

	resources := make(ResourceCollection, 0, len(collection))

	for _, model := range collection {
		resources = append(resources, FormatResource(model))
	}

	return ResourceList{
		Integrations: resources,
	}
}
