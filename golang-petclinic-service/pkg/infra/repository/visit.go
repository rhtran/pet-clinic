package repository

import (
	"github.com/rhtran/golang-petclinic-service/pkg/model"
	"time"
)

// Visit
// represents a visit made by a pet to the veterinarian.
// entity:model visit
type Visit struct {
	model.Base
	VisitDate   time.Time
	Description string
	PetID       int
	Pet         Pet
}
