package test

import (
	"github.com/jinzhu/gorm"
	"github.com/manyminds/api2go"
	"log"
	"time"
)

type mockedContext struct {
	db *gorm.DB
}

// NewMockedContext mocks an `api2go.APIContexter` to be used in tests.
func NewMockedContext() api2go.APIContexter {
	return &mockedContext{}
}

func (*mockedContext) Deadline() (deadline time.Time, ok bool) {
	return time.Now(), false
}

func (*mockedContext) Done() <-chan struct{} {
	return make(chan struct{})
}

func (*mockedContext) Err() error {
	return nil
}

func (*mockedContext) Value(key interface{}) interface{} {
	return nil
}

func (*mockedContext) Set(key string, value interface{}) {
}

func (ctx *mockedContext) Get(key string) (interface{}, bool) {
	switch key {
	case "db":
		if ctx.db != nil {
			return ctx.db, true
		}

		db, err := NewTestDatabase()
		if err != nil {
			log.Fatal(err)
		}

		ctx.db = db

		return db, true
	}

	return nil, false
}

func (*mockedContext) Reset() {
}
