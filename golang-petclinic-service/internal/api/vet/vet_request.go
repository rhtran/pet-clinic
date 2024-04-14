package vet

import (
	validation "github.com/go-ozzo/ozzo-validation/v3"
	"github.com/rhtran/golang-petclinic-service/pkg/infra/repository/vet"
	model2 "github.com/rhtran/golang-petclinic-service/pkg/model"
)

type Request struct {
	ID          int             `json:"-"`
	FirstName   string          `json:"firstName" binding:"required"`
	LastName    string          `json:"lastName" binding:"required"`
	Specialties []vet.Specialty `json:"specialties" binding:"required"`
}

func (vr Request) Validate() error {
	return validation.ValidateStruct(&vr,
		validation.Field(&vr.FirstName, validation.Required),
		validation.Field(&vr.LastName, validation.Required),
		validation.Field(&vr.Specialties, validation.Required),
	)
}

func (vr Request) ToVet(vetRequest *Request) *vet.Vet {
	//specialties := make([]Specialty, len(vetRequest.Specialties))

	//for i, v := range vetRequest.Specialties {
	//	specialties[i] = Specialty{
	//		Base: model.Base{ID: vetRequest.Specialties[i].ID},
	//		Name: vetRequest.Specialties[i].Name,
	//	}
	//}

	return &vet.Vet{
		Base: model2.Base{
			ID: vetRequest.ID,
		},
		Person: model2.Person{
			FirstName: vetRequest.FirstName,
			LastName:  vetRequest.LastName,
		},
		Specialties: vetRequest.Specialties,
	}
}
