package model

import (
	"github.com/satori/go.uuid"
	"testing"
	"time"
)

func TestModel_GetID(t *testing.T) {
	baseID, _ := uuid.FromString("01234567-89ab-cdef-0123-456789abcdef")

	type fields struct {
		ID        uuid.UUID
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt *time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"base", fields{ID: baseID}, "01234567-89ab-cdef-0123-456789abcdef"},
		{"empty", fields{}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Model{
				ID:        tt.fields.ID,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
			}
			if got := m.GetID(); got != tt.want {
				t.Errorf("Model.GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModel_SetID(t *testing.T) {
	type fields struct {
		ID        uuid.UUID
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt *time.Time
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"base", fields{}, args{"01234567-89ab-cdef-0123-456789abcdef"}, false},
		{"capitalized", fields{}, args{"01234567-89AB-CDEF-0123-456789ABCDEF"}, false},
		{"short", fields{}, args{"01234567-89ab-cdef-0123-456789abcde"}, true},
		{"long", fields{}, args{"01234567-89ab-cdef-0123-456789abcdef0"}, true},
		{"invalid-character", fields{}, args{"01234567-89ab-cdef-g123-456789abcdef"}, true},
		{"empty", fields{}, args{""}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Model{
				ID:        tt.fields.ID,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
			}
			if err := m.SetID(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Model.SetID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestModel_BeforeCreate(t *testing.T) {
	baseID, _ := uuid.FromString("01234567-89ab-cdef-0123-456789abcdef")

	type fields struct {
		ID        uuid.UUID
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt *time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"base", fields{baseID, time.Time{}, time.Time{}, nil}, false},
		{"deleted", fields{baseID, time.Time{}, time.Time{}, &time.Time{}}, true},
		{"empty", fields{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Model{
				ID:        tt.fields.ID,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
			}
			if err := m.BeforeCreate(); (err != nil) != tt.wantErr {
				t.Errorf("Model.BeforeCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
