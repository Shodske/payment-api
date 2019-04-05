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

func TestOrganisationSource_Create(t *testing.T) {
	req := NewMockedRequest()

	baseOrg := &model.Organisation{
		Name: "Base Organisation",
	}
	idOrg := &model.Organisation{
		Model: model.Model{
			ID: uuid.NewV4(),
		},
		Name: "Organisation with ID",
	}

	baseRes := &api2go.Response{
		Code: http.StatusCreated,
		Res:  baseOrg,
	}
	idRes := &api2go.Response{
		Code: http.StatusCreated,
		Res:  idOrg,
	}

	type args struct {
		obj interface{}
		req api2go.Request
	}
	tests := []struct {
		name    string
		src     *OrganisationSource
		args    args
		want    api2go.Responder
		wantErr bool
	}{
		{"base", &OrganisationSource{}, args{baseOrg, *req}, baseRes, false},
		{"with-id", &OrganisationSource{}, args{idOrg, *req}, idRes, false},
		{"duplicate-id", &OrganisationSource{}, args{idOrg, *req}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			src := &OrganisationSource{}
			got, err := src.Create(tt.args.obj, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrganisationSource.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrganisationSource.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrganisationSource_FindAll(t *testing.T) {
	req := NewMockedRequest()

	baseRes := &api2go.Response{
		Code: http.StatusOK,
		Res:  GetOrganisationFixtures(false),
	}

	type args struct {
		req api2go.Request
	}
	tests := []struct {
		name    string
		src     *OrganisationSource
		args    args
		want    api2go.Responder
		wantErr bool
	}{
		{"base", &OrganisationSource{}, args{*req}, baseRes, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			src := &OrganisationSource{}
			got, err := src.FindAll(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrganisationSource.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrganisationSource.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrganisationSource_PaginatedFindAll(t *testing.T) {
	req := NewMockedRequest()
	firstReq := NewMockedRequest()
	firstReq.QueryParams = map[string][]string{"page[number]": {"1"}, "page[size]": {"1"}}
	secondReq := NewMockedRequest()
	secondReq.QueryParams = map[string][]string{"page[number]": {"2"}, "page[size]": {"1"}}
	oorReq := NewMockedRequest()
	oorReq.QueryParams = map[string][]string{"page[number]": {"100"}, "page[size]": {"100"}}

	orgs := GetOrganisationFixtures(false)
	count := len(orgs)

	firstRes := &api2go.Response{
		Code: http.StatusOK,
		Res:  orgs[0:1],
	}
	secondRes := &api2go.Response{
		Code: http.StatusOK,
		Res:  orgs[1:2],
	}
	emptyRes := &api2go.Response{
		Code: http.StatusOK,
		Res:  make([]*model.Organisation, 0),
	}

	type args struct {
		req api2go.Request
	}
	tests := []struct {
		name           string
		src            *OrganisationSource
		args           args
		wantTotalCount uint
		wantResponse   api2go.Responder
		wantErr        bool
	}{
		{"first-page", &OrganisationSource{}, args{*firstReq}, uint(count), firstRes, false},
		{"second-page", &OrganisationSource{}, args{*secondReq}, uint(count), secondRes, false},
		{"out-of-range", &OrganisationSource{}, args{*oorReq}, uint(count), emptyRes, false},
		{"not-paginated", &OrganisationSource{}, args{*req}, 0, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			src := &OrganisationSource{}
			gotTotalCount, gotResponse, err := src.PaginatedFindAll(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrganisationSource.PaginatedFindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTotalCount != tt.wantTotalCount {
				t.Errorf("OrganisationSource.PaginatedFindAll() gotTotalCount = %v, want %v", gotTotalCount, tt.wantTotalCount)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("OrganisationSource.PaginatedFindAll() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestOrganisationSource_FindOne(t *testing.T) {
	req := NewMockedRequest()

	orgs := GetOrganisationFixtures(false)
	deletedOrgs := GetOrganisationFixtures(true)

	type args struct {
		ID  string
		req api2go.Request
	}
	type testData struct {
		name    string
		src     *OrganisationSource
		args    args
		want    api2go.Responder
		wantErr bool
	}
	tests := []testData{
		{"deleted", &OrganisationSource{}, args{deletedOrgs[0].GetID(), *req}, nil, true},
	}

	for i, org := range orgs {
		res := &api2go.Response{
			Code: http.StatusOK,
			Res:  org,
		}
		tests = append(
			tests,
			testData{"organisation-" + string(i), &OrganisationSource{}, args{org.GetID(), *req}, res, false},
		)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			src := &OrganisationSource{}
			got, err := src.FindOne(tt.args.ID, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrganisationSource.FindOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrganisationSource.FindOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrganisationSource_Update(t *testing.T) {
	req := NewMockedRequest()
	org := GetOrganisationFixtures(false)[0]
	deletedOrg := GetOrganisationFixtures(true)[0]

	updateData := &model.Organisation{
		Model: model.Model{ID: org.ID},
		Name:  "Updated Organisation",
	}
	updatedOrg := *org
	updatedOrg.Name = "Updated Organisation"

	res := &api2go.Response{
		Code: http.StatusOK,
		Res:  &updatedOrg,
	}

	noIDData := &model.Organisation{
		Name: "Updated Organisation",
	}

	delUpdateData := &model.Organisation{
		Model: model.Model{ID: deletedOrg.ID},
		Name:  "Updated Organisation",
	}

	type args struct {
		obj interface{}
		req api2go.Request
	}
	tests := []struct {
		name    string
		src     *OrganisationSource
		args    args
		want    api2go.Responder
		wantErr bool
	}{
		{"base", &OrganisationSource{}, args{updateData, *req}, res, false},
		{"no-id", &OrganisationSource{}, args{noIDData, *req}, nil, true},
		{"deleted", &OrganisationSource{}, args{delUpdateData, *req}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			src := &OrganisationSource{}
			got, err := src.Update(tt.args.obj, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrganisationSource.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil && tt.want != nil {
				// Updated at should be changed.
				gotRes := got.Result().(*model.Organisation)
				wantRes := tt.want.Result().(*model.Organisation)
				if reflect.DeepEqual(gotRes.UpdatedAt, wantRes.UpdatedAt) {
					t.Errorf("Organisation.UpdatedAt should be updated, got %v, want %v", gotRes.UpdatedAt, wantRes.UpdatedAt)
				}

				// set updated at to an empty time, so the next compare won't fail.
				gotRes.UpdatedAt = time.Time{}
				wantRes.UpdatedAt = time.Time{}
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrganisationSource.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrganisationSource_Delete(t *testing.T) {
	req := NewMockedRequest()

	orgs := GetOrganisationFixtures(false)
	deletedOrgs := GetOrganisationFixtures(true)

	type args struct {
		id  string
		req api2go.Request
	}

	type testData struct {
		name    string
		src     *OrganisationSource
		args    args
		want    api2go.Responder
		wantErr bool
	}

	tests := []testData{
		{"invalid", &OrganisationSource{}, args{"not-a-uuid", *req}, nil, true},
		{"deleted", &OrganisationSource{}, args{deletedOrgs[0].GetID(), *req}, nil, true},
	}

	for i, org := range orgs {
		res := &api2go.Response{
			Code: http.StatusNoContent,
		}
		tests = append(
			tests,
			testData{"organisation-" + string(i), &OrganisationSource{}, args{org.GetID(), *req}, res, false},
		)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			src := &OrganisationSource{}
			got, err := src.Delete(tt.args.id, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrganisationSource.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrganisationSource.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
