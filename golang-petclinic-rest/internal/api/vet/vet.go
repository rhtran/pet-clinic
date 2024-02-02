package vet

import (
	model2 "github.com/rhtran/golang-petclinic-rest/pkg/model"
)

// many2many relation by join vet_specialty table
type Vet struct {
	model2.Base
	model2.Person
	Specialties []Specialty `gorm:"many2many:vet_specialties;",json:"specialties"`
}

type Specialty struct {
	model2.Base
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
