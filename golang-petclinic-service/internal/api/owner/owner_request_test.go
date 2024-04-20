package owner

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToOwner(t *testing.T) {
	ownerRequest := &Request{
		FirstName: "John",
		LastName:  "Doe",
		Username:  "jdoe",
		Address:   "123 Main St",
		City:      "Anytown",
		Telephone: "1234567890",
	}

	owner := ToOwner(ownerRequest)

	assert.Equal(t, ownerRequest.FirstName, owner.Person.FirstName)
	assert.Equal(t, ownerRequest.LastName, owner.Person.LastName)
	assert.Equal(t, ownerRequest.Username, owner.Username)
	assert.Equal(t, ownerRequest.Address, owner.Address)
	assert.Equal(t, ownerRequest.City, owner.City)
	assert.Equal(t, ownerRequest.Telephone, owner.Telephone)
}

func TestToOwnerWithEmptyFields(t *testing.T) {
	ownerRequest := &Request{
		FirstName: "",
		LastName:  "",
		Username:  "",
		Address:   "",
		City:      "",
		Telephone: "",
	}

	owner := ToOwner(ownerRequest)

	assert.Equal(t, ownerRequest.FirstName, owner.Person.FirstName)
	assert.Equal(t, ownerRequest.LastName, owner.Person.LastName)
	assert.Equal(t, ownerRequest.Username, owner.Username)
	assert.Equal(t, ownerRequest.Address, owner.Address)
	assert.Equal(t, ownerRequest.City, owner.City)
	assert.Equal(t, ownerRequest.Telephone, owner.Telephone)
}

func TestToOwnerWithPartialFields(t *testing.T) {
	ownerRequest := &Request{
		FirstName: "John",
		LastName:  "Doe",
		Username:  "",
		Address:   "",
		City:      "",
		Telephone: "",
	}

	owner := ToOwner(ownerRequest)

	assert.Equal(t, ownerRequest.FirstName, owner.Person.FirstName)
	assert.Equal(t, ownerRequest.LastName, owner.Person.LastName)
	assert.Equal(t, ownerRequest.Username, owner.Username)
	assert.Equal(t, ownerRequest.Address, owner.Address)
	assert.Equal(t, ownerRequest.City, owner.City)
	assert.Equal(t, ownerRequest.Telephone, owner.Telephone)
}
