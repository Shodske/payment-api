package source

import (
	"errors"
	"github.com/Shodske/payment-api/pkg/model"
	"github.com/manyminds/api2go"
	"net/http"
)

// PaymentSource struct that implements the different interfaces for handling CRUD actions on Payment Models.
type PaymentSource struct {
}

// Create method required to implement `api2go.ResourceCreator`. Implementing this interface will enable the URI:
// POST /payments
func (src *PaymentSource) Create(obj interface{}, req api2go.Request) (api2go.Responder, error) {
	payment, ok := obj.(*model.Payment)
	if !ok {
		return nil, api2go.NewHTTPError(errors.New("invalid type"), "invalid type", http.StatusConflict)
	}

	db, err := getDatabase(req)
	if err != nil {
		return nil, err
	}

	if err := db.Create(payment).Error; err != nil {
		return nil, err
	}

	return &api2go.Response{Res: payment, Code: http.StatusCreated}, nil
}

// FindAll method required to implement `api2go.FindAll`. Implementing this interface will enable the URI:
// GET /payments
func (src *PaymentSource) FindAll(req api2go.Request) (api2go.Responder, error) {
	db, err := getDatabase(req)
	if err != nil {
		return nil, err
	}

	payments := make([]*model.Payment, 0)
	if err := db.Find(&payments).Error; err != nil {
		return nil, err
	}

	return &api2go.Response{Res: payments, Code: http.StatusOK}, nil
}

// PaginatedFindAll method required to implement `api2go.PaginatedFindAll`. Implementing this interface will enable the URI:
// GET /payments?page[number]=<number>&page[size]=<size>
func (src *PaymentSource) PaginatedFindAll(req api2go.Request) (uint, api2go.Responder, error) {
	number, size, err := extractPaginationQuery(req)
	if err != nil {
		return 0, nil, err
	}

	db, err := getDatabase(req)
	if err != nil {
		return 0, nil, err
	}

	var count uint
	db.Model(&model.Payment{}).Count(&count)

	payments := make([]*model.Payment, 0)
	db.Limit(size).Offset((number - 1) * size).Find(&payments)

	return count, &api2go.Response{Res: payments, Code: http.StatusOK}, nil
}

// FindOne method required to implement `api2go.ResourceGetter`. Implementing this interface will enable the URI:
// GET /payments/:paymentID
func (src *PaymentSource) FindOne(id string, req api2go.Request) (api2go.Responder, error) {
	db, err := getDatabase(req)
	if err != nil {
		return nil, err
	}

	payment := &model.Payment{}
	if err := payment.SetID(id); err != nil {
		return nil, api2go.NewHTTPError(err, "invalid id", http.StatusBadRequest)
	}

	if err := db.Where(payment).First(payment).Error; err != nil {
		return nil, api2go.NewHTTPError(err, "could not find payments resource", http.StatusNotFound)
	}

	return &api2go.Response{Res: payment, Code: http.StatusOK}, nil
}

// Update method required to implement `api2go.ResourceUpdater`. Implementing this interface will enable the URI:
// PATCH /payments/:paymentID
func (src *PaymentSource) Update(obj interface{}, req api2go.Request) (api2go.Responder, error) {
	paymentData, ok := obj.(*model.Payment)
	if !ok {
		return nil, api2go.NewHTTPError(errors.New("invalid type"), "invalid type", http.StatusConflict)
	}

	if paymentData.GetID() == "" {
		return nil, api2go.NewHTTPError(errors.New("missing id"), "missing id", http.StatusConflict)
	}

	db, err := getDatabase(req)
	if err != nil {
		return nil, err
	}

	payment := &model.Payment{Model: model.Model{ID: paymentData.ID}}
	if err := db.Where(payment).First(payment).Error; err != nil {
		return nil, err
	}
	if err := db.Model(payment).Update(paymentData).Error; err != nil {
		return nil, err
	}

	return &api2go.Response{Res: payment, Code: http.StatusOK}, nil
}

// Delete method required to implement `api2go.ResourceDeleter`. Implementing this interface will enable the URI:
// DELETE /payments/:paymentID
func (src *PaymentSource) Delete(id string, req api2go.Request) (api2go.Responder, error) {
	if id == "" {
		return nil, api2go.NewHTTPError(errors.New("invalid id"), "invalid id", http.StatusBadRequest)
	}

	db, err := getDatabase(req)
	if err != nil {
		return nil, err
	}

	payment := &model.Payment{}
	if err := payment.SetID(id); err != nil {
		return nil, api2go.NewHTTPError(err, "invalid id", http.StatusBadRequest)
	}

	if err := db.Where(payment).First(payment).Error; err != nil {
		return nil, api2go.NewHTTPError(err, "could not find payments resource", http.StatusNotFound)
	}

	if err := db.Delete(payment).Error; err != nil {
		return nil, err
	}

	return &api2go.Response{Code: http.StatusNoContent}, nil
}
