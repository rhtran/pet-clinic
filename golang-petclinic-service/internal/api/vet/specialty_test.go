package vet

import (
	"github.com/rhtran/golang-petclinic-service/pkg/infra/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestFromSpecialty(t *testing.T) {
	specialty := &repository.Specialty{Model: gorm.Model{ID: 1}, Name: "Cardiology"}
	s := &Specialty{}
	s.FromSpecialty(specialty)

	assert.Equal(t, specialty.ID, s.ID)
	assert.Equal(t, specialty.Name, s.Name)
}

func TestFromSpecialties(t *testing.T) {
	specialties := []repository.Specialty{
		{Model: gorm.Model{ID: 1}, Name: "Cardiology"},
		{Model: gorm.Model{ID: 2}, Name: "Dentistry"},
		{Model: gorm.Model{ID: 3}, Name: "Empty Specialty"},
	}

	result := FromSpecialties(specialties)

	assert.Equal(t, len(specialties), len(*result))
	for i, specialty := range *result {
		assert.Equal(t, specialties[i].ID, specialty.ID)
		assert.Equal(t, specialties[i].Name, specialty.Name)
	}
}

func TestFromSpecialtiesWithEmptySlice(t *testing.T) {
	specialties := []repository.Specialty{}

	result := FromSpecialties(specialties)

	assert.Equal(t, 0, len(*result))
}

func TestToSpecialties(t *testing.T) {
	specialties := []Specialty{
		{ID: 1, Name: "Cardiology"},
		{ID: 2, Name: "Dentistry"},
		{ID: 3, Name: "Empty Specialty"},
	}

	result := ToSpecialties(specialties)

	assert.Equal(t, len(specialties), len(*result))
	for i, specialty := range *result {
		assert.Equal(t, specialties[i].ID, specialty.ID)
		assert.Equal(t, specialties[i].Name, specialty.Name)
	}
}

func TestToSpecialtiesWithEmptySlice(t *testing.T) {
	specialties := []Specialty{}

	result := ToSpecialties(specialties)

	assert.Equal(t, 0, len(*result))
}
