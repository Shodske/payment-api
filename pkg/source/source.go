package source

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/manyminds/api2go"
	"net/http"
	"strconv"
)

// Extract the page number and size from the request.
func extractPaginationQuery(req api2go.Request) (number int64, size int64, err error) {
	numberQuery, ok := req.QueryParams["page[number]"]
	if !ok {
		return 0, 0, api2go.NewHTTPError(
			errors.New("could not find `page[number]` in query"),
			"could not find `page[number]` in query",
			http.StatusBadRequest,
		)
	}

	number, err = strconv.ParseInt(numberQuery[0], 10, 64)
	if err != nil {
		return 0, 0, api2go.NewHTTPError(
			err,
			"invalid value for `page[number]` in query",
			http.StatusBadRequest,
		)
	}

	sizeQuery, ok := req.QueryParams["page[size]"]
	if !ok {
		return 0, 0, api2go.NewHTTPError(
			errors.New("could not find `page[size]` in query"),
			"could not find `page[size]` in query",
			http.StatusBadRequest,
		)
	}

	size, err = strconv.ParseInt(sizeQuery[0], 10, 64)
	if err != nil {
		return 0, 0, api2go.NewHTTPError(
			err,
			"invalid value for `page[size]` in query",
			http.StatusBadRequest,
		)
	}

	return
}

// Get the gorm database from the request's context.
func getDatabase(req api2go.Request) (*gorm.DB, error) {
	db, ok := req.Context.Get("db")
	if !ok {
		return nil, errors.New("no database connection")
	}

	conn, ok := db.(*gorm.DB)
	if !ok {
		return nil, errors.New("error accessing database")
	}

	return conn, nil
}
