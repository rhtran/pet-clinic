package vet

import (
	validation "github.com/go-ozzo/ozzo-validation/v3"
	"wire-petclinic-service/pkg/model"
)

type VetRequest struct {
	ID          int         `json:"-"`
	FirstName   string      `json:"firstName" binding:"required"`
	LastName    string      `json:"lastName" binding:"required"`
	Specialties []Specialty `json:"specialties" binding:"required"`
}

func (vr VetRequest) Validate() error {
	return validation.ValidateStruct(&vr,
		validation.Field(&vr.FirstName, validation.Required),
		validation.Field(&vr.LastName, validation.Required),
		validation.Field(&vr.Specialties, validation.Required),
	)
}

func (vr VetRequest) ToVet(vetRequest *VetRequest) *Vet {
	//specialties := make([]Specialty, len(vetRequest.Specialties))

	//for i, v := range vetRequest.Specialties {
	//	specialties[i] = Specialty{
	//		Base: model.Base{ID: vetRequest.Specialties[i].ID},
	//		Name: vetRequest.Specialties[i].Name,
	//	}
	//}

	return &Vet{
		Base: model.Base{
			ID: vetRequest.ID,
		},
		Person: model.Person{
			FirstName: vetRequest.FirstName,
			LastName:  vetRequest.LastName,
		},
		Specialties: vetRequest.Specialties,
	}
}
