package vet

import (
	model2 "github.com/rhtran/golang-petclinic-service/pkg/model"
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
