package vet

import (
	"github.com/rhtran/golang-petclinic-service/pkg/infra/repository"
	"gorm.io/gorm"
)

type Specialty struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (s *Specialty) FromSpecialty(specialty *repository.Specialty) {
	s.ID = specialty.ID
	s.Name = specialty.Name
}

func FromSpecialties(specialties []repository.Specialty) *[]Specialty {
	specialtyResponses := make([]Specialty, len(specialties))
	for i, s := range specialties {
		specialtyResponses[i].ID = s.ID
		specialtyResponses[i].Name = s.Name
	}
	return &specialtyResponses
}

func ToSpecialty(specialty *Specialty) *repository.Specialty {
	return &repository.Specialty{
		Model: gorm.Model{
			ID: specialty.ID,
		},
		Name: specialty.Name,
	}
}

func ToSpecialties(specialties []Specialty) *[]repository.Specialty {
	specialtyEntities := make([]repository.Specialty, len(specialties))
	for i, s := range specialties {
		specialtyEntities[i] = *ToSpecialty(&s)
	}

	return &specialtyEntities
}
