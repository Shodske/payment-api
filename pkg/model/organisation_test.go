package model

import (
	"reflect"
	"testing"

	"github.com/manyminds/api2go/jsonapi"
)

func TestOrganisation_GetReferences(t *testing.T) {
	type fields struct {
		Model Model
		Name  string
	}
	tests := []struct {
		name   string
		fields fields
		want   []jsonapi.Reference
	}{
		{
			"base-case",
			fields{},
			[]jsonapi.Reference{
				{
					Name:         "payments",
					Type:         "payments",
					IsNotLoaded:  true,
					Relationship: jsonapi.ToManyRelationship,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			org := &Organisation{
				Model: tt.fields.Model,
				Name:  tt.fields.Name,
			}
			if got := org.GetReferences(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Organisation.GetReferences() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrganisation_GetReferencedIDs(t *testing.T) {
	type fields struct {
		Model Model
		Name  string
	}
	tests := []struct {
		name   string
		fields fields
		want   []jsonapi.ReferenceID
	}{
		{"base-case", fields{}, []jsonapi.ReferenceID{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			org := &Organisation{
				Model: tt.fields.Model,
				Name:  tt.fields.Name,
			}
			if got := org.GetReferencedIDs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Organisation.GetReferencedIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}
