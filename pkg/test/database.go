package test

import (
	"github.com/Shodske/payment-api/pkg/model"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"log"
	"time"
)

var organisationFixtures = []*model.Organisation{
	{
		Model: model.Model{ID: uuid.NewV4()},
		Name:  "Test Organisation A",
	},
	{
		Model: model.Model{ID: uuid.NewV4()},
		Name:  "Test Organisation B",
	},
	{
		Model: model.Model{ID: uuid.NewV4(), DeletedAt: &time.Time{}},
		Name:  "Deleted Test Organisation",
	},
}

// GetOrganisationFixtures method returns all Organisation fixtures used to populate the test database. When `deleted`
// bool is set to true, it returns all Organisations marked as deleted, all other Organisations otherwise.
func GetOrganisationFixtures(deleted bool) []*model.Organisation {
	orgs := []*model.Organisation{}

	for _, org := range organisationFixtures {
		if (org.DeletedAt != nil) == deleted {
			orgs = append(orgs, org)
		}
	}

	return orgs
}

// NewTestDatabase method creates a database connection for testing purposes, using sqlite.
func NewTestDatabase() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "/tmp/api.db")
	if err != nil {
		log.Fatal(err)
	}

	if err = migrate(db); err != nil {
		return nil, err
	}

	if err = populate(db); err != nil {
		return nil, err
	}

	return db, nil
}

// Create all database tables.
func migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.Organisation{},
		&model.Party{},
		&model.Charge{},
		&model.FX{},
		&model.Payment{},
	).Error
}

// Populate the database with test data.
func populate(db *gorm.DB) error {
	for _, org := range organisationFixtures {
		if err := db.Create(org).Error; err != nil {
			return err
		}
	}

	return nil
}
