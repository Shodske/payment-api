package main

import (
	"fmt"
	"github.com/Shodske/payment-api/pkg/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

func main() {
	log.Print("starting payment-api server")

	// Open a database connection to auto migrate the schemas.
	log.Print("migrating schemas...")
	conn, err := openDatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}

	conn.AutoMigrate(
		&model.Organisation{},
		&model.Party{},
		&model.Charge{},
		&model.FX{},
		&model.Payment{},
	)

	// Close the connection after we're done.
	if err = conn.Close(); err != nil {
		log.Fatal(err)
	}
}

// Open a `gorm.DB` database connection, which can be used to migrate the database tables and execute CRUD actions on
// Models.
func openDatabaseConnection() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	conn, err := gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
			host,
			port,
			database,
			username,
			password,
		),
	)

	if conn != nil {
		// TODO Set LogMode based on the current log level, which should be set through an env variable.
		conn.LogMode(true)
	}

	return conn, err
}
