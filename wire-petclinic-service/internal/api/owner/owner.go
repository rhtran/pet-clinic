package owner

import (
	"wire-petclinic-service/internal/api/pet"
	"wire-petclinic-service/pkg/model"
)

/*
Owner has many Pets, OwnerID is the foreign key
Has Many configuration
*/
type Owner struct {
	model.Base
	model.Person
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	Address   string    `json:"address"`
	City      string    `json:"city"`
	Telephone string    `json:"telephone"`
	Pets      []pet.Pet `json:"pets"`
}

func (o *Owner) ToOwnerResponse(owner *Owner) *OwnerResponse {
	petP := &pet.Pet{}
	var petResponses []pet.PetResponse

	if len(owner.Pets) > 0 {
		petResponses = petP.ToPetResponses(owner.Pets)
	}

	return &OwnerResponse{
		ID:        owner.ID,
		Username:  owner.Username,
		FirstName: owner.FirstName,
		LastName:  owner.LastName,
		Address:   owner.Address,
		City:      owner.City,
		Telephone: owner.Telephone,
		Pets:      petResponses,
	}
}

func (o *Owner) ToOwnerResponses(owners []Owner) []OwnerResponse {
	ownerResponses := make([]OwnerResponse, len(owners))
	for i, v := range owners {
		ownerResponses[i] = *o.ToOwnerResponse(&v)
	}
	return ownerResponses
}
