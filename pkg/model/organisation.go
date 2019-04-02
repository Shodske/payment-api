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

// Method required to implement `jsonapi.MarshalReferences`.
func (org *Organisation) GetReferences() []jsonapi.Reference {
	panic("implement me")
}

// Method required to implement `jsonapi.MarshalLinkedRelations`.
func (org *Organisation) GetReferencedIDs() []jsonapi.ReferenceID {
	panic("implement me")
}
