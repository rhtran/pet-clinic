package visit

import (
	"time"
	"wire-petclinic-service/internal/api/pet"
)

type VisitResponse struct {
	ID          int             `json:"id"`
	VisitDate   time.Time       `json:"visitDate"`
	Description string          `json:"description"`
	PetResponse pet.PetResponse `json:"pet"`
}
