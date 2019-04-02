# Payment API
This is an example repo for a small RESTful API in which you can manage
payments.

This API is based on the [json:api](https://jsonapi.org/) specification
and one can expect the API to behave according to this specification.

There are two resources defined in this API:
- `organisations`
- `payments`

## Organisations
Organisations are small resource that only have a name and serve as a
relationship for payments.

### Endpoints
- `GET /v0/organisations`
- `GET /v0/organisations/:organisationID`
- `POST /v0/organisations`
- `PATCH /v0/organisations/:organisationID`
- `DELETE /v0/organisations/:organisationID`

### Example Resource
```json
{
  "type": "organisations",
  "id": "e5dbc976-5d51-487e-a414-c1ca517ee6bc",
  "attributes": {
    "name": "string"
  }
}
```

## Payments
Payments resources hold all the data required for a payment. They have a
relationship with organisations.

### Endpoints
- `GET /v0/payments`
- `GET /v0/payments/:paymentID`
- `POST /v0/payments`
- `PATCH /v0/payments/:paymentID`
- `DELETE /v0/payments/:paymentID`

### Example Resource
```json
{
  "type": "payments",
  "id": "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43",
  "attributes": {
    "amount": "100.21",
    "currency": "GBP",
    "end_to_end_reference": "Wil piano Jan",
    "numeric_reference": "1002001",
    "payment_id": "123456789012345678",
    "payment_purpose": "Paying for goods/services",
    "payment_scheme": "FPS",
    "payment_type": "Credit",
    "processing_date": "2017-01-18",
    "reference": "Payment for Em's piano lessons",
    "scheme_payment_sub_type": "InternetBanking",
    "scheme_payment_type": "ImmediatePayment",
    
    "beneficiary_party": {
      "account_name": "W Owens",
      "account_number": "31926819",
      "account_number_code": "BBAN",
      "account_type": 0,
      "address": "1 The Beneficiary Localtown SE2",
      "bank_id": "403000",
      "bank_id_code": "GBDSC",
      "name": "Wilfred Jeremiah Owens"
    },
    
    "debtor_party": {
      "account_name": "EJ Brown Black",
      "account_number": "GB29XABC10161234567801",
      "account_number_code": "IBAN",
      "address": "10 Debtor Crescent Sourcetown NE1",
      "bank_id": "203301",
      "bank_id_code": "GBDSC",
      "name": "Emelia Jane Brown"
    },
    
    "sponsor_party": {
      "account_number": "56781234",
      "bank_id": "123123",
      "bank_id_code": "GBDSC"
    },
    
    "charges_information": {
      "bearer_code": "SHAR",
      "sender_charges": [
        {
          "amount": "5.00",
          "currency": "GBP"
        },
        {
          "amount": "10.00",
          "currency": "USD"
        }
      ],
      "receiver_charges_amount": "1.00",
      "receiver_charges_currency": "USD"
    },
    
    "fx": {
      "contract_reference": "FX123",
      "exchange_rate": "2.00000",
      "original_amount": "200.42",
      "original_currency": "USD"
    }
  },
  "relationships": {
    "organisation": {
      "data": {
        "type": "organisations",
        "id": "e5dbc976-5d51-487e-a414-c1ca517ee6bc"
      }
    }
  }
}
```

## Running the API
A `docker-compose.yml` file has been set up to start the API. Which can
be run using the `Makefile`. To override any of the environment
variables, you can create a `.env` file in the project root, which will
be picked up by `docker-compose`.

- `make start`: Start the Docker containers running the API and the
                Postgres database.
- `make stop`: Stop all Docker containers.
- `make logs`: Follow the logs of the API and Postgres database.
- `make test`: Run all tests.
