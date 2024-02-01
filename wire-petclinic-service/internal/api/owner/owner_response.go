package owner

import (
	"wire-petclinic-service/internal/api/pet"
)

type OwnerResponse struct {
	ID        int               `json:"id"`
	FirstName string            `json:"firstName"`
	LastName  string            `json:"lastName"`
	Username  string            `json:"username"`
	Address   string            `json:"address"`
	City      string            `json:"city"`
	Telephone string            `json:"telephone"`
	Pets      []pet.PetResponse `json:"pets,omitempty"`
}
