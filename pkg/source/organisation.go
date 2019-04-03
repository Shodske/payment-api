package source

import (
	"errors"
	"github.com/Shodske/payment-api/pkg/model"
	"github.com/manyminds/api2go"
	"net/http"
)

// OrganisationSource struct that implements the different interfaces for handling CRUD actions on Organisation Models.
type OrganisationSource struct {
}

// Create method required to implement `api2go.ResourceCreator`. Implementing this interface will enable the URI:
// POST /organisations
func (src *OrganisationSource) Create(obj interface{}, req api2go.Request) (api2go.Responder, error) {
	org, ok := obj.(*model.Organisation)
	if !ok {
		return nil, api2go.NewHTTPError(errors.New("invalid type"), "invalid type", http.StatusConflict)
	}

	db, err := getDatabase(req)
	if err != nil {
		return nil, err
	}

	if err := db.Create(org).Error; err != nil {
		return nil, err
	}

	return &api2go.Response{Res: org, Code: http.StatusCreated}, nil
}

// FindAll method required to implement `api2go.FindAll`. Implementing this interface will enable the URI:
// GET /organisations
func (src *OrganisationSource) FindAll(req api2go.Request) (api2go.Responder, error) {
	db, err := getDatabase(req)
	if err != nil {
		return nil, err
	}

	orgs := make([]*model.Organisation, 0)
	db.Find(&orgs)

	return &api2go.Response{Res: orgs, Code: http.StatusOK}, nil
}

// PaginatedFindAll method required to implement `api2go.PaginatedFindAll`. Implementing this interface will enable the URI:
// GET /organisations?page[number]=<number>&page[size]=<size>
func (src *OrganisationSource) PaginatedFindAll(req api2go.Request) (uint, api2go.Responder, error) {
	number, size, err := extractPaginationQuery(req)
	if err != nil {
		return 0, nil, err
	}

	db, err := getDatabase(req)
	if err != nil {
		return 0, nil, err
	}

	var count uint
	db.Model(&model.Organisation{}).Count(&count)

	orgs := make([]*model.Organisation, 0)
	db.Limit(size).Offset((number - 1) * size).Find(&orgs)

	return count, &api2go.Response{Res: orgs, Code: http.StatusOK}, nil
}

// FindOne method required to implement `api2go.ResourceGetter`. Implementing this interface will enable the URI:
// GET /organisations/:organisationID
func (src *OrganisationSource) FindOne(id string, req api2go.Request) (api2go.Responder, error) {
	db, err := getDatabase(req)
	if err != nil {
		return nil, err
	}

	org := &model.Organisation{}
	if err := org.SetID(id); err != nil {
		return nil, api2go.NewHTTPError(err, "invalid id", http.StatusBadRequest)
	}

	if err := db.Where(org).First(org).Error; err != nil {
		return nil, api2go.NewHTTPError(err, "could not find organisations resource", http.StatusNotFound)
	}

	return &api2go.Response{Res: org, Code: http.StatusOK}, nil
}

// Update method required to implement `api2go.ResourceUpdater`. Implementing this interface will enable the URI:
// PATCH /organisations/:organisationID
func (src *OrganisationSource) Update(obj interface{}, req api2go.Request) (api2go.Responder, error) {
	orgData, ok := obj.(*model.Organisation)
	if !ok {
		return nil, api2go.NewHTTPError(errors.New("invalid type"), "invalid type", http.StatusConflict)
	}

	if orgData.GetID() == "" {
		return nil, api2go.NewHTTPError(errors.New("missing id"), "missing id", http.StatusConflict)
	}

	db, err := getDatabase(req)
	if err != nil {
		return nil, err
	}

	org := &model.Organisation{Model: model.Model{ID: orgData.ID}}
	if err := db.Where(org).First(org).Error; err != nil {
		return nil, err
	}
	if err := db.Model(org).Update(orgData).Error; err != nil {
		return nil, err
	}

	return &api2go.Response{Res: org, Code: http.StatusOK}, nil
}

// Delete method required to implement `api2go.ResourceDeleter`. Implementing this interface will enable the URI:
// DELETE /organisations/:organisationID
func (src *OrganisationSource) Delete(id string, req api2go.Request) (api2go.Responder, error) {
	if id == "" {
		return nil, api2go.NewHTTPError(errors.New("invalid id"), "invalid id", http.StatusBadRequest)
	}

	db, err := getDatabase(req)
	if err != nil {
		return nil, err
	}

	org := &model.Organisation{}
	if err := org.SetID(id); err != nil {
		return nil, api2go.NewHTTPError(err, "invalid id", http.StatusBadRequest)
	}

	if err := db.Where(org).First(org).Error; err != nil {
		return nil, api2go.NewHTTPError(err, "could not find organisations resource", http.StatusNotFound)
	}

	if err := db.Delete(org).Error; err != nil {
		return nil, err
	}

	return &api2go.Response{Code: http.StatusNoContent}, nil
}
