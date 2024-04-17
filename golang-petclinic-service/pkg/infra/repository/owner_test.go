package repository

import (
	"github.com/rhtran/golang-petclinic-service/pkg/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOwnerFields(t *testing.T) {
	owner := &Owner{
		Person: model.Person{
			FirstName: "John",
			LastName:  "Doe",
		},
		Username:  "johndoe",
		Password:  "password",
		Address:   "123 Main St",
		City:      "Anytown",
		Telephone: "1234567890",
		Pets:      []Pet{},
	}

	assert.Equal(t, "John", owner.FirstName)
	assert.Equal(t, "Doe", owner.LastName)
	assert.Equal(t, "johndoe", owner.Username)
	assert.Equal(t, "password", owner.Password)
	assert.Equal(t, "123 Main St", owner.Address)
	assert.Equal(t, "Anytown", owner.City)
	assert.Equal(t, "1234567890", owner.Telephone)
	assert.Empty(t, owner.Pets)
}

func TestOwnerFieldsWithPets(t *testing.T) {
	pet := Pet{
		Name: "Fido",
		Type: Type{
			Name: "Dog",
		},
	}

	owner := &Owner{
		Person: model.Person{
			FirstName: "John",
			LastName:  "Doe",
		},
		Username:  "johndoe",
		Password:  "password",
		Address:   "123 Main St",
		City:      "Anytown",
		Telephone: "1234567890",
		Pets:      []Pet{pet},
	}

	assert.Equal(t, "John", owner.FirstName)
	assert.Equal(t, "Doe", owner.LastName)
	assert.Equal(t, "johndoe", owner.Username)
	assert.Equal(t, "password", owner.Password)
	assert.Equal(t, "123 Main St", owner.Address)
	assert.Equal(t, "Anytown", owner.City)
	assert.Equal(t, "1234567890", owner.Telephone)
	assert.NotEmpty(t, owner.Pets)
	assert.Equal(t, "Fido", owner.Pets[0].Name)
	assert.Equal(t, "Dog", owner.Pets[0].Type.Name)
}
