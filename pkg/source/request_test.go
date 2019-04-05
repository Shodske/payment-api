package source

import "github.com/manyminds/api2go"

// NewMockedRequest mocks an `api2go.Request` to be used in tests.
func NewMockedRequest() *api2go.Request {
	return &api2go.Request{
		Context: NewMockedContext(),
	}
}
