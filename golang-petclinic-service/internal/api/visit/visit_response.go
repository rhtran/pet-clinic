package visit

import (
	"github.com/rhtran/golang-petclinic-service/internal/api/pet"
	"github.com/rhtran/golang-petclinic-service/pkg/infra/repository"
	"time"
)

type Response struct {
	ID          int          `json:"id"`
	VisitDate   time.Time    `json:"visitDate"`
	Description string       `json:"description"`
	PetResponse pet.Response `json:"pet"`
}

func (vr *Response) FromVisit(visit *repository.Visit) {
	vr.ID = visit.ID
	vr.VisitDate = visit.VisitDate
	vr.Description = visit.Description
	vr.PetResponse = pet.Response{}
	vr.PetResponse.FromPet(&visit.Pet)
}

func FromVisits(visits []repository.Visit) []Response {
	visitResponses := make([]Response, len(visits))
	for i, v := range visits {
		visitResponses[i].FromVisit(&v)
	}
	return visitResponses
}
