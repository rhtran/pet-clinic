package vet

import (
	"github.com/rhtran/golang-petclinic-service/pkg/infra/repository"
	"github.com/rhtran/golang-petclinic-service/pkg/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestResponseFromVet(t *testing.T) {
	vet := &repository.Vet{
		Model: gorm.Model{ID: 1},
		Person: model.Person{
			FirstName: "John",
			LastName:  "Doe",
		},
		Specialties: []repository.Specialty{
			{Model: gorm.Model{ID: 1}, Name: "Cardiology"},
			{Model: gorm.Model{ID: 2}, Name: "Dentistry"},
		},
	}
	response := &Response{}
	response.FromVet(vet)

	assert.Equal(t, vet.ID, response.ID)
	assert.Equal(t, vet.FirstName, response.FirstName)
	assert.Equal(t, vet.LastName, response.LastName)
	assert.Equal(t, len(vet.Specialties), len(response.Specialties))
	for i, specialty := range response.Specialties {
		assert.Equal(t, vet.Specialties[i].ID, specialty.ID)
		assert.Equal(t, vet.Specialties[i].Name, specialty.Name)
	}
}

func TestFromVets(t *testing.T) {
	vets := []repository.Vet{
		{
			Model: gorm.Model{ID: 1},
			Person: model.Person{
				FirstName: "John",
				LastName:  "Doe",
			},
			Specialties: []repository.Specialty{
				{Model: gorm.Model{ID: 1}, Name: "Cardiology"},
				{Model: gorm.Model{ID: 2}, Name: "Dentistry"},
			},
		},
		{
			Model: gorm.Model{ID: 2},
			Person: model.Person{
				FirstName: "John",
				LastName:  "Doe",
			},
			Specialties: []repository.Specialty{
				{Model: gorm.Model{ID: 3}, Name: "Neurology"},
			},
		},
	}

	responses := FromVets(vets)

	assert.Equal(t, len(vets), len(responses))
	for i, response := range responses {
		assert.Equal(t, vets[i].ID, response.ID)
		assert.Equal(t, vets[i].FirstName, response.FirstName)
		assert.Equal(t, vets[i].LastName, response.LastName)
		assert.Equal(t, len(vets[i].Specialties), len(response.Specialties))
		for j, specialty := range response.Specialties {
			assert.Equal(t, vets[i].Specialties[j].ID, specialty.ID)
			assert.Equal(t, vets[i].Specialties[j].Name, specialty.Name)
		}
	}
}

func TestFromVetsWithEmptySlice(t *testing.T) {
	vets := []repository.Vet{}

	responses := FromVets(vets)

	assert.Equal(t, 0, len(responses))
}
