package visit

import (
	"github.com/rhtran/golang-petclinic-rest/internal/api/pet"
	"github.com/rhtran/golang-petclinic-rest/pkg/model"
	"time"
)

type Visit struct {
	model.Base
	VisitDate   time.Time
	Description string
	PetID       int
	Pet         pet.Pet
}

func (v *Visit) ToVisitResponse(visit *Visit) *VisitResponse {
	return &VisitResponse{
		ID:          visit.ID,
		VisitDate:   visit.VisitDate,
		Description: visit.Description,
		PetResponse: *visit.Pet.ToPetResponse(&visit.Pet),
	}
}

func (v Visit) ToVisitResponses(visits []Visit) []VisitResponse {
	visitResponses := make([]VisitResponse, len(visits))
	for i, v := range visits {
		visitResponses[i] = *v.ToVisitResponse(&v)
	}
	return visitResponses
}
