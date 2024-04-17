package repository

import (
	"github.com/rhtran/golang-petclinic-service/pkg/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVetSpecialties(t *testing.T) {
	vet := &Vet{
		Person: model.Person{
			FirstName: "John",
			LastName:  "Doe",
		},
		Specialties: []Specialty{
			{Name: "Cardiology"},
			{Name: "Dentistry"},
		},
	}

	assert.Equal(t, "Cardiology", vet.Specialties[0].Name)
	assert.Equal(t, "Dentistry", vet.Specialties[1].Name)
}

func TestVetSpecialtiesWithEmptyName(t *testing.T) {
	vet := &Vet{
		Person: model.Person{
			FirstName: "John",
			LastName:  "Doe",
		},
		Specialties: []Specialty{
			{Name: ""},
		},
	}

	assert.Equal(t, "", vet.Specialties[0].Name)
}

func TestVetSpecialtiesWithNoSpecialties(t *testing.T) {
	vet := &Vet{
		Person: model.Person{
			FirstName: "John",
			LastName:  "Doe",
		},
		Specialties: []Specialty{},
	}

	assert.Empty(t, vet.Specialties)
}

func TestSpecialtyName(t *testing.T) {
	specialty := &Specialty{
		Name: "Cardiology",
	}

	assert.Equal(t, "Cardiology", specialty.Name)
}

func TestSpecialtyNameWithEmptyName(t *testing.T) {
	specialty := &Specialty{
		Name: "",
	}

	assert.Equal(t, "", specialty.Name)
}
