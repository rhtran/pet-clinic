package visit

import (
	"github.com/rhtran/golang-petclinic-rest/internal/api/pet"
	"time"
)

type VisitResponse struct {
	ID          int             `json:"id"`
	VisitDate   time.Time       `json:"visitDate"`
	Description string          `json:"description"`
	PetResponse pet.PetResponse `json:"pet"`
}
