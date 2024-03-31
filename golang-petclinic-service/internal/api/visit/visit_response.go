package visit

import (
	"github.com/rhtran/golang-petclinic-service/internal/api/pet"
	"time"
)

type Response struct {
	ID          int          `json:"id"`
	VisitDate   time.Time    `json:"visitDate"`
	Description string       `json:"description"`
	PetResponse pet.Response `json:"pet"`
}
