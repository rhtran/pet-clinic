package repository

import (
	"github.com/rhtran/golang-petclinic-service/pkg/model"
	"gorm.io/gorm"
)

// many2many relation by join vet_specialty table
type Vet struct {
	gorm.Model
	model.Person
	Specialties []Specialty `gorm:"many2many:vet_specialties;",json:"specialties"`
}

type Specialty struct {
	gorm.Model
	Name string
}
