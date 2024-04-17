package vet

import (
	validation "github.com/go-ozzo/ozzo-validation/v3"
	"github.com/rhtran/golang-petclinic-service/pkg/infra/repository"
	model2 "github.com/rhtran/golang-petclinic-service/pkg/model"
	"gorm.io/gorm"
)

type Request struct {
	ID          uint        `json:"-"`
	FirstName   string      `json:"firstName" binding:"required"`
	LastName    string      `json:"lastName" binding:"required"`
	Specialties []Specialty `json:"specialties" binding:"required"`
}

func (vr Request) Validate() error {
	return validation.ValidateStruct(&vr,
		validation.Field(&vr.FirstName, validation.Required),
		validation.Field(&vr.LastName, validation.Required),
		validation.Field(&vr.Specialties, validation.Required),
	)
}

func ToVet(vetRequest *Request) *repository.Vet {
	return &repository.Vet{
		Model: gorm.Model{
			ID: vetRequest.ID,
		},
		Person: model2.Person{
			FirstName: vetRequest.FirstName,
			LastName:  vetRequest.LastName,
		},
		Specialties: *ToSpecialties(vetRequest.Specialties),
	}
}
