[![Go Report Card](https://goreportcard.com/badge/github.com/Shodske/payment-api)](https://goreportcard.com/report/github.com/Shodske/payment-api)

# Payment API
This is an example repo for a small RESTful API in which you can manage
payments.

This API is based on the [json:api](https://jsonapi.org/) specification
and one can expect the API to behave according to this specification.

For ease of use a `docker-compose.yml` has been created that will run
the API, database and Swagger UI.

## API Specification
The API specification is written in [OpenAPI Specification 3.0](https://swagger.io/specification/)
and can found in `api/oas.yml`.

A Swagger UI docker image is set up for easier reading and sending
requests to the API.

## Running the API
A `docker-compose.yml` file has been set up to start the API. Which can
be run using the `Makefile`. To override any of the environment
variables, you can create a `.env` file in the project root, which will
be picked up by `docker-compose`.

- `make start`: Start the Docker containers running the API, the Postgres
                database and Swagger UI.
- `make stop`: Stop all Docker containers.
- `make logs`: Follow the logs of the API and Postgres database.
- `make test`: Run all tests.
