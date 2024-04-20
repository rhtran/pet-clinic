package owner

import (
	"github.com/rhtran/golang-petclinic-service/internal/api/pet"
	"github.com/rhtran/golang-petclinic-service/pkg/infra/repository"
	"github.com/rhtran/golang-petclinic-service/pkg/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestResponseFromOwner(t *testing.T) {
	owner := &repository.Owner{
		Model: gorm.Model{
			ID: 1,
		},
		Person: model.Person{
			FirstName: "John",
			LastName:  "Doe",
		},
		Username:  "jdoe",
		Address:   "123 Main St",
		City:      "Anytown",
		Telephone: "1234567890",
		Pets:      []repository.Pet{},
	}

	response := &Response{}
	response.FromOwner(owner)

	assert.Equal(t, owner.ID, response.ID)
	assert.Equal(t, owner.FirstName, response.FirstName)
	assert.Equal(t, owner.LastName, response.LastName)
	assert.Equal(t, owner.Username, response.Username)
	assert.Equal(t, owner.Address, response.Address)
	assert.Equal(t, owner.City, response.City)
	assert.Equal(t, owner.Telephone, response.Telephone)
	assert.Equal(t, pet.FromPets(owner.Pets), response.Pets)
}

func TestFromOwners(t *testing.T) {
	owners := []repository.Owner{
		{
			Model: gorm.Model{ID: 1},
			Person: model.Person{
				FirstName: "John",
				LastName:  "Doe",
			},
			Username:  "jdoe",
			Address:   "123 Main St",
			City:      "Anytown",
			Telephone: "1234567890",
			Pets:      []repository.Pet{},
		},
		{
			Model: gorm.Model{ID: 2},
			Person: model.Person{FirstName: "Jane",
				LastName: "Doe"},
			Username:  "jadoe",
			Address:   "456 Main St",
			City:      "Anytown",
			Telephone: "0987654321",
			Pets:      []repository.Pet{},
		},
	}

	responses := FromOwners(owners)

	assert.Equal(t, len(owners), len(responses))
	for i, owner := range owners {
		assert.Equal(t, owner.ID, responses[i].ID)
		assert.Equal(t, owner.FirstName, responses[i].FirstName)
		assert.Equal(t, owner.LastName, responses[i].LastName)
		assert.Equal(t, owner.Username, responses[i].Username)
		assert.Equal(t, owner.Address, responses[i].Address)
		assert.Equal(t, owner.City, responses[i].City)
		assert.Equal(t, owner.Telephone, responses[i].Telephone)
		assert.Equal(t, pet.FromPets(owner.Pets), responses[i].Pets)
	}
}
