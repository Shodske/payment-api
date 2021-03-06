package main

import (
	"context"
	"fmt"
	"github.com/Shodske/payment-api/pkg/model"
	"github.com/Shodske/payment-api/pkg/source"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/manyminds/api2go"
	"log"
	"net/http"
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
		&model.CurrencyAmount{},
		&model.FX{},
		&model.Payment{},
	)

	// Close the connection after we're done.
	if err = conn.Close(); err != nil {
		log.Fatal(err)
	}

	log.Print("initialising api...")
	api := initAPI()

	port := os.Getenv("PORT")
	log.Printf("server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, api.Handler()))
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
		conn = conn.Set("gorm:auto_preload", true)
	}

	return conn, err
}

// Initialise the API with required middleware and registered resources.
func initAPI() *api2go.API {
	api := api2go.NewAPI("v0")

	// Allow cross origin requests
	api.UseMiddleware(func(_ api2go.APIContexter, res http.ResponseWriter, _ *http.Request) {
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
		res.Header().Set("Access-Control-Allow-Headers", "*")
	})

	// For every request open a database connection.
	api.UseMiddleware(func(ctx api2go.APIContexter, _ http.ResponseWriter, req *http.Request) {
		db, err := openDatabaseConnection()
		if err != nil {
			return
		}
		ctx.Set("db", db)

		// Whenever the request is done, close the database connection.
		go func(conn *gorm.DB, ctx context.Context) {
			<-ctx.Done()
			conn.Close()
		}(db, req.Context())
	})

	api.AddResource(&model.Organisation{}, &source.OrganisationSource{})
	api.AddResource(&model.Payment{}, &source.PaymentSource{})

	return api
}
