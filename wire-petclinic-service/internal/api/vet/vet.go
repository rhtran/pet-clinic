package vet

import (
	"wire-petclinic-service/pkg/model"
)

// many2many relation by join vet_specialty table
type Vet struct {
	model.Base
	model.Person
	Specialties []Specialty `gorm:"many2many:vet_specialties;",json:"specialties"`
}

type Specialty struct {
	model.Base
	Name string `json:"name"`
}

func (v *Vet) ToVetResponse(vet *Vet) *VetResponse {
	return &VetResponse{
		ID:          vet.ID,
		FirstName:   vet.FirstName,
		LastName:    vet.LastName,
		Specialties: vet.Specialties,
	}
}

func (v Vet) ToVetResponses(vets []Vet) []VetResponse {
	vetResponses := make([]VetResponse, len(vets))
	for i, v := range vets {
		vetResponses[i] = *v.ToVetResponse(&v)
	}
	return vetResponses
}
