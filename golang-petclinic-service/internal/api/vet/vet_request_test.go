package vet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRequestValidation(t *testing.T) {
	vetRequest := &Request{
		FirstName: "John",
		LastName:  "Doe",
		Specialties: []Specialty{
			{ID: 1, Name: "Cardiology"},
			{ID: 2, Name: "Dentistry"},
		},
	}

	err := vetRequest.Validate()

	assert.Nil(t, err)
}

func TestRequestValidationWithEmptyFields(t *testing.T) {
	vetRequest := &Request{
		FirstName:   "",
		LastName:    "",
		Specialties: []Specialty{},
	}

	err := vetRequest.Validate()

	assert.NotNil(t, err)
}

func TestToVet(t *testing.T) {
	vetRequest := &Request{
		FirstName: "John",
		LastName:  "Doe",
		Specialties: []Specialty{
			{ID: 1, Name: "Cardiology"},
			{ID: 2, Name: "Dentistry"},
		},
	}

	vet := ToVet(vetRequest)

	assert.Equal(t, vetRequest.FirstName, vet.FirstName)
	assert.Equal(t, vetRequest.LastName, vet.LastName)
	assert.Equal(t, len(vetRequest.Specialties), len(vet.Specialties))
	for i, specialty := range vet.Specialties {
		assert.Equal(t, vetRequest.Specialties[i].ID, specialty.ID)
		assert.Equal(t, vetRequest.Specialties[i].Name, specialty.Name)
	}
}
