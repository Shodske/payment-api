package test

import (
	"github.com/Shodske/payment-api/pkg/model"
	"github.com/jinzhu/gorm"
	"github.com/manyminds/api2go"
	"log"
	"time"
)

type mockedContext struct {
}

// NewMockedContext mocks an `api2go.APIContexter` to be used in tests.
func NewMockedContext() api2go.APIContexter {
	return &mockedContext{}
}

func (mockedContext) Deadline() (deadline time.Time, ok bool) {
	return time.Now(), false
}

func (mockedContext) Done() <-chan struct{} {
	return make(chan struct{})
}

func (mockedContext) Err() error {
	return nil
}

func (mockedContext) Value(key interface{}) interface{} {
	return nil
}

func (mockedContext) Set(key string, value interface{}) {
}

func (mockedContext) Get(key string) (interface{}, bool) {
	switch key {
	case "db":
		db, err := gorm.Open("sqlite3", "/tmp/api.db")
		if err != nil {
			log.Fatal(err)
		}

		db.AutoMigrate(
			&model.Organisation{},
			&model.Party{},
			&model.Charge{},
			&model.FX{},
			&model.Payment{},
		)

		return db, true
	}

	return nil, false
}

func (mockedContext) Reset() {
}
