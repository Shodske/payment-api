package model

import (
	"errors"
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
	if id == "" {
		return nil
	}

	ID, err := uuid.FromString(id)
	if err != nil {
		return err
	}

	m.ID = ID

	return nil
}

// Method required to implement `jsonapi.MarshalIdentifier`. Returns a string representation of the uuid of the Model.
func (m *Model) GetID() string {
	// An empty uuid yields a string representation of only zeroes.
	if id := m.ID.String(); id != "00000000-0000-0000-0000-000000000000" {
		return id
	}

	return ""
}

// Callback method for `gorm` package. Makes sure a uuid is set before a Model is created.
func (m *Model) BeforeCreate() error {
	if m.DeletedAt != nil {
		return errors.New("cannot create deleted Model")
	}
	if m.GetID() == "" {
		m.ID = uuid.NewV4()
	}
	return nil
}
