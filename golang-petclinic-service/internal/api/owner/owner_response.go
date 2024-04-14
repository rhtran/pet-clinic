package owner

import (
	"github.com/rhtran/golang-petclinic-service/infra/repository/owner"
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

func (or *Response) FromOwner(owner *owner.Owner) {
	or.ID = owner.ID
	or.Username = owner.Username
	or.FirstName = owner.FirstName
	or.LastName = owner.LastName
	or.Address = owner.Address
	or.City = owner.City
	or.Telephone = owner.Telephone
	or.Pets = pet.FromPets(owner.Pets)
}

func FromOwners(owners []owner.Owner) []Response {
	ownerResponses := make([]Response, len(owners))
	for i, v := range owners {
		ownerResponses[i].FromOwner(&v)
	}
	return ownerResponses
}
