package owner

import (
	"github.com/rhtran/golang-petclinic-service/internal/api/pet"
)

type Response struct {
	ID        int            `json:"id"`
	FirstName string         `json:"firstName"`
	LastName  string         `json:"lastName"`
	Username  string         `json:"username"`
	Address   string         `json:"address"`
	City      string         `json:"city"`
	Telephone string         `json:"telephone"`
	Pets      []pet.Response `json:"pets,omitempty"`
}
