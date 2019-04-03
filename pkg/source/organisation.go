package source

import (
	"errors"
	"github.com/Shodske/payment-api/pkg/model"
	"github.com/jinzhu/gorm"
	"github.com/manyminds/api2go"
	"net/http"
	"strconv"
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

	db, ok := req.Context.Get("db")
	if !ok {
		return nil, errors.New("no database connection")
	}

	conn, ok := db.(*gorm.DB)
	if !ok {
		return nil, errors.New("error accessing database")
	}

	if err := conn.Create(org).Error; err != nil {
		return nil, err
	}

	return &api2go.Response{Res: org, Code: http.StatusCreated}, nil
}

// FindAll method required to implement `api2go.FindAll`. Implementing this interface will enable the URI:
// GET /organisations
func (src *OrganisationSource) FindAll(req api2go.Request) (api2go.Responder, error) {
	db, ok := req.Context.Get("db")
	if !ok {
		return nil, errors.New("no database connection")
	}

	conn, ok := db.(*gorm.DB)
	if !ok {
		return nil, errors.New("error accessing database")
	}

	orgs := make([]*model.Organisation, 0)
	conn.Find(&orgs)

	return &api2go.Response{Res: orgs, Code: http.StatusOK}, nil
}

// PaginatedFindAll method required to implement `api2go.PaginatedFindAll`. Implementing this interface will enable the URI:
// GET /organisations?page[number]=<number>&page[size]=<size>
func (src *OrganisationSource) PaginatedFindAll(req api2go.Request) (uint, api2go.Responder, error) {
	numberQuery, ok := req.QueryParams["page[number]"]
	if !ok {
		return 0, nil, api2go.NewHTTPError(
			errors.New("could not find `page[number]` in query"),
			"could not find `page[number]` in query",
			http.StatusBadRequest,
		)
	}

	number, err := strconv.ParseInt(numberQuery[0], 10, 64)
	if err != nil {
		return 0, nil, api2go.NewHTTPError(
			err,
			"invalid value for `page[number]` in query",
			http.StatusBadRequest,
		)
	}

	sizeQuery, ok := req.QueryParams["page[size]"]
	if !ok {
		return 0, nil, api2go.NewHTTPError(
			errors.New("could not find `page[size]` in query"),
			"could not find `page[size]` in query",
			http.StatusBadRequest,
		)
	}

	size, err := strconv.ParseInt(sizeQuery[0], 10, 64)
	if err != nil {
		return 0, nil, api2go.NewHTTPError(
			err,
			"invalid value for `page[size]` in query",
			http.StatusBadRequest,
		)
	}

	db, ok := req.Context.Get("db")
	if !ok {
		return 0, nil, errors.New("no database connection")
	}

	conn, ok := db.(*gorm.DB)
	if !ok {
		return 0, nil, errors.New("error accessing database")
	}

	var count uint
	conn.Model(&model.Organisation{}).Count(&count)

	orgs := make([]*model.Organisation, 0)
	conn.Limit(size).Offset((number - 1) * size).Find(&orgs)

	return count, &api2go.Response{Res: orgs, Code: http.StatusOK}, nil
}

// FindOne method required to implement `api2go.ResourceGetter`. Implementing this interface will enable the URI:
// GET /organisations/:organisationID
func (src *OrganisationSource) FindOne(id string, req api2go.Request) (api2go.Responder, error) {
	db, ok := req.Context.Get("db")
	if !ok {
		return nil, errors.New("no database connection")
	}

	conn, ok := db.(*gorm.DB)
	if !ok {
		return nil, errors.New("error accessing database")
	}

	org := &model.Organisation{}
	if err := org.SetID(id); err != nil {
		return nil, api2go.NewHTTPError(err, "invalid id", http.StatusBadRequest)
	}

	if err := conn.Where(org).First(org).Error; err != nil {
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

	db, ok := req.Context.Get("db")
	if !ok {
		return nil, errors.New("no database connection")
	}

	conn, ok := db.(*gorm.DB)
	if !ok {
		return nil, errors.New("error accessing database")
	}

	org := &model.Organisation{Model: model.Model{ID: orgData.ID}}
	if err := conn.Where(org).First(org).Error; err != nil {
		return nil, err
	}
	if err := conn.Model(org).Update(orgData).Error; err != nil {
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

	db, ok := req.Context.Get("db")
	if !ok {
		return nil, errors.New("no database connection")
	}

	conn, ok := db.(*gorm.DB)
	if !ok {
		return nil, errors.New("error accessing database")
	}

	org := &model.Organisation{}
	if err := org.SetID(id); err != nil {
		return nil, api2go.NewHTTPError(err, "invalid id", http.StatusBadRequest)
	}

	if err := conn.Where(org).First(org).Error; err != nil {
		return nil, api2go.NewHTTPError(err, "could not find organisations resource", http.StatusNotFound)
	}

	if err := conn.Delete(org).Error; err != nil {
		return nil, err
	}

	return &api2go.Response{Code: http.StatusNoContent}, nil
}
