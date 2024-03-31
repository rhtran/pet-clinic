package visit

import (
	"github.com/rhtran/golang-petclinic-service/internal/api/pet"
	"github.com/rhtran/golang-petclinic-service/pkg/model"
	"time"
)

type Visit struct {
	model.Base
	VisitDate   time.Time
	Description string
	PetID       int
	Pet         pet.Pet
}

func (v *Visit) ToVisitResponse(visit *Visit) *Response {
	return &Response{
		ID:          visit.ID,
		VisitDate:   visit.VisitDate,
		Description: visit.Description,
		PetResponse: *visit.Pet.ToPetResponse(&visit.Pet),
	}
}

func (v Visit) ToVisitResponses(visits []Visit) []Response {
	visitResponses := make([]Response, len(visits))
	for i, v := range visits {
		visitResponses[i] = *v.ToVisitResponse(&v)
	}
	return visitResponses
}
