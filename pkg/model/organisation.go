package model

import (
	"github.com/manyminds/api2go/jsonapi"
)

// Organisation model that represents an organisation. Can be marshaled to a json resource according to the json:api
// specification.
type Organisation struct {
	Model `json:"-"`
	Name  string `json:"name"`
}

// GetReferences method required to implement `jsonapi.MarshalReferences`.
func (org *Organisation) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Name:         "payments",
			Type:         "payments",
			IsNotLoaded:  true,
			Relationship: jsonapi.ToManyRelationship,
		},
	}
}

// GetReferencedIDs method required to implement `jsonapi.MarshalLinkedRelations`.
func (org *Organisation) GetReferencedIDs() []jsonapi.ReferenceID {
	// Always return an empty slice. We don't want to retrieve all payments for an organisation, as they can be quite
	// numerous.
	return []jsonapi.ReferenceID{}
}
