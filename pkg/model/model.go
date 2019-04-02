package model

import (
	"github.com/satori/go.uuid"
	"time"
)

// Base struct for models that have a uuid as primary key. Based on `gorm.Model`, to be embedded in other models.
// Implements `` and
type Model struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// Method required to implement `jsonapi.UnmarshalIdentifier`. Converts the string representation of a uuid to the
// internal Model representation.
func (m *Model) SetID(id string) error {
	panic("implement me")
}

// Method required to implement `jsonapi.MarshalIdentifier`. Returns a string representation of the uuid of the Model.
func (m *Model) GetID() string {
	panic("implement me")
}

// Callback method for `gorm` package. Makes sure a uuid is set before a Model is created.
func (m *Model) BeforeCreate() error {
	panic("implement me")
}
