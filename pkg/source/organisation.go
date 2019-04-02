package source

import (
	"github.com/manyminds/api2go"
)

// OrganisationSource struct that implements the different interfaces for handling CRUD actions on Organisation Models.
type OrganisationSource struct {
}

// Create method required to implement `api2go.ResourceCreator`. Implementing this interface will enable the URI:
// POST /organisations
func (src *OrganisationSource) Create(obj interface{}, req api2go.Request) (api2go.Responder, error) {
	panic("implement me")
}

// FindAll method required to implement `api2go.FindAll`. Implementing this interface will enable the URI:
// GET /organisations
func (src *OrganisationSource) FindAll(req api2go.Request) (api2go.Responder, error) {
	panic("implement me")
}

// PaginatedFindAll method required to implement `api2go.PaginatedFindAll`. Implementing this interface will enable the URI:
// GET /organisations?page[number]=<number>&page[size]=<size>
func (src *OrganisationSource) PaginatedFindAll(req api2go.Request) (uint, api2go.Responder, error) {
	panic("implement me")
}

// FindOne method required to implement `api2go.ResourceGetter`. Implementing this interface will enable the URI:
// GET /organisations/:organisationID
func (src *OrganisationSource) FindOne(id string, req api2go.Request) (api2go.Responder, error) {
	panic("implement me")
}

// Update method required to implement `api2go.ResourceUpdater`. Implementing this interface will enable the URI:
// PATCH /organisations/:organisationID
func (src *OrganisationSource) Update(obj interface{}, req api2go.Request) (api2go.Responder, error) {
	panic("implement me")
}

// Delete method required to implement `api2go.ResourceDeleter`. Implementing this interface will enable the URI:
// DELETE /organisations/:organisationID
func (src *OrganisationSource) Delete(id string, req api2go.Request) (api2go.Responder, error) {
	panic("implement me")
}
