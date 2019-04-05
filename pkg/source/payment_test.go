package source

import (
	"github.com/Shodske/payment-api/pkg/model"
	"github.com/manyminds/api2go"
	"github.com/satori/go.uuid"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestPaymentSource_Create(t *testing.T) {
	req := NewMockedRequest()
	payments := GetPaymentFixtures(false)

	// For ease of use, we copy a payment from the fixtures and use that.
	basePayment := *payments[0]
	basePayment.ID = uuid.UUID{}
	basePayment.Organisation = model.Organisation{}

	idPayment := basePayment
	idPayment.ID = uuid.NewV4()

	baseRes := &api2go.Response{
		Code: http.StatusCreated,
		Res:  &basePayment,
	}
	idRes := &api2go.Response{
		Code: http.StatusCreated,
		Res:  &idPayment,
	}

	type args struct {
		obj interface{}
		req api2go.Request
	}
	tests := []struct {
		name    string
		src     *PaymentSource
		args    args
		want    api2go.Responder
		wantErr bool
	}{
		{"base", &PaymentSource{}, args{&basePayment, *req}, baseRes, false},
		{"with-id", &PaymentSource{}, args{&idPayment, *req}, idRes, false},
		{"duplicate-id", &PaymentSource{}, args{&idPayment, *req}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			src := &PaymentSource{}
			got, err := src.Create(tt.args.obj, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("PaymentSource.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentSource.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaymentSource_FindAll(t *testing.T) {
	req := NewMockedRequest()

	baseRes := &api2go.Response{
		Code: http.StatusOK,
		Res:  GetPaymentFixtures(false),
	}

	type args struct {
		req api2go.Request
	}
	tests := []struct {
		name    string
		src     *PaymentSource
		args    args
		want    api2go.Responder
		wantErr bool
	}{
		{"base", &PaymentSource{}, args{*req}, baseRes, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			src := &PaymentSource{}
			got, err := src.FindAll(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("PaymentSource.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentSource.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaymentSource_PaginatedFindAll(t *testing.T) {
	req := NewMockedRequest()
	firstReq := NewMockedRequest()
	firstReq.QueryParams = map[string][]string{"page[number]": {"1"}, "page[size]": {"2"}}
	secondReq := NewMockedRequest()
	secondReq.QueryParams = map[string][]string{"page[number]": {"2"}, "page[size]": {"2"}}
	oorReq := NewMockedRequest()
	oorReq.QueryParams = map[string][]string{"page[number]": {"100"}, "page[size]": {"100"}}

	payments := GetPaymentFixtures(false)
	count := len(payments)

	firstRes := &api2go.Response{
		Code: http.StatusOK,
		Res:  payments[0:2],
	}
	secondRes := &api2go.Response{
		Code: http.StatusOK,
		Res:  payments[2:3],
	}
	emptyRes := &api2go.Response{
		Code: http.StatusOK,
		Res:  make([]*model.Payment, 0),
	}

	type args struct {
		req api2go.Request
	}
	tests := []struct {
		name    string
		src     *PaymentSource
		args    args
		want    uint
		want1   api2go.Responder
		wantErr bool
	}{
		{"first-page", &PaymentSource{}, args{*firstReq}, uint(count), firstRes, false},
		{"second-page", &PaymentSource{}, args{*secondReq}, uint(count), secondRes, false},
		{"out-of-range", &PaymentSource{}, args{*oorReq}, uint(count), emptyRes, false},
		{"not-paginated", &PaymentSource{}, args{*req}, 0, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			src := &PaymentSource{}
			got, got1, err := src.PaginatedFindAll(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("PaymentSource.PaginatedFindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PaymentSource.PaginatedFindAll() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PaymentSource.PaginatedFindAll() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPaymentSource_FindOne(t *testing.T) {
	req := NewMockedRequest()

	payments := GetPaymentFixtures(false)
	deletedPayments := GetPaymentFixtures(true)

	type args struct {
		id  string
		req api2go.Request
	}
	type testData struct {
		name    string
		src     *PaymentSource
		args    args
		want    api2go.Responder
		wantErr bool
	}
	tests := []testData{
		{"deleted", &PaymentSource{}, args{deletedPayments[0].GetID(), *req}, nil, true},
	}

	for i, payment := range payments {
		res := &api2go.Response{
			Code: http.StatusOK,
			Res:  payment,
		}
		tests = append(
			tests,
			testData{"payments-" + string(i), &PaymentSource{}, args{payment.GetID(), *req}, res, false},
		)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			src := &PaymentSource{}
			got, err := src.FindOne(tt.args.id, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("PaymentSource.FindOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentSource.FindOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaymentSource_Update(t *testing.T) {
	req := NewMockedRequest()
	payment := GetPaymentFixtures(false)[0]
	deletedOrg := GetPaymentFixtures(true)[0]

	updateData := &model.Payment{
		Model:     model.Model{ID: payment.ID},
		Reference: "Updated Payment",
	}
	updatedPayemnt := *payment
	updatedPayemnt.Reference = "Updated Payment"

	res := &api2go.Response{
		Code: http.StatusOK,
		Res:  &updatedPayemnt,
	}

	noIDData := &model.Payment{
		Reference: "Updated Payment",
	}

	delUpdateData := &model.Payment{
		Model:     model.Model{ID: deletedOrg.ID},
		Reference: "Updated Payment",
	}

	type args struct {
		obj interface{}
		req api2go.Request
	}
	tests := []struct {
		name    string
		src     *PaymentSource
		args    args
		want    api2go.Responder
		wantErr bool
	}{
		{"base", &PaymentSource{}, args{updateData, *req}, res, false},
		{"no-id", &PaymentSource{}, args{noIDData, *req}, nil, true},
		{"deleted", &PaymentSource{}, args{delUpdateData, *req}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			src := &PaymentSource{}
			got, err := src.Update(tt.args.obj, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("PaymentSource.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil && tt.want != nil {
				// Updated at should be changed.
				gotRes := got.Result().(*model.Payment)
				wantRes := tt.want.Result().(*model.Payment)
				if reflect.DeepEqual(gotRes.UpdatedAt, wantRes.UpdatedAt) {
					t.Errorf("Payment.UpdatedAt should be updated, got %v, want %v", gotRes.UpdatedAt, wantRes.UpdatedAt)
				}

				// set updated at to an empty time, so the next compare won't fail.
				gotRes.UpdatedAt = time.Time{}
				wantRes.UpdatedAt = time.Time{}
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentSource.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaymentSource_Delete(t *testing.T) {
	req := NewMockedRequest()

	payments := GetPaymentFixtures(false)
	deletedPayments := GetPaymentFixtures(true)

	type args struct {
		id  string
		req api2go.Request
	}

	type testData struct {
		name    string
		src     *PaymentSource
		args    args
		want    api2go.Responder
		wantErr bool
	}

	tests := []testData{
		{"invalid", &PaymentSource{}, args{"not-a-uuid", *req}, nil, true},
		{"deleted", &PaymentSource{}, args{deletedPayments[0].GetID(), *req}, nil, true},
	}

	for i, payment := range payments {
		res := &api2go.Response{
			Code: http.StatusNoContent,
		}
		tests = append(
			tests,
			testData{"payment-" + string(i), &PaymentSource{}, args{payment.GetID(), *req}, res, false},
		)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			src := &PaymentSource{}
			got, err := src.Delete(tt.args.id, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("PaymentSource.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentSource.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
