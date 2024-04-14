package vet

import (
	"github.com/rhtran/golang-petclinic-service/pkg/infra/repository"
)

type Response struct {
	ID          int                    `json:"id"`
	FirstName   string                 `json:"firstName"`
	LastName    string                 `json:"lastName"`
	Specialties []repository.Specialty `json:"specialties,omitempty"`
}

type Responses struct {
	Vets []Response `json:"vets"`
}

func (vr *Response) FromVet(vet *repository.Vet) {
	vr.ID = vet.ID
	vr.FirstName = vet.FirstName
	vr.LastName = vet.LastName
	vr.Specialties = vet.Specialties
}

func FromVets(vets []repository.Vet) []Response {
	vetResponses := make([]Response, len(vets))
	for i, v := range vets {
		vetResponses[i].FromVet(&v)
	}
	return vetResponses
}
