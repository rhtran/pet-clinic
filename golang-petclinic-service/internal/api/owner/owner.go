package owner

import (
	pet "github.com/rhtran/golang-petclinic-service/internal/api/pet"
	model "github.com/rhtran/golang-petclinic-service/pkg/model"
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

func (o *Owner) ToOwnerResponse(owner *Owner) *Response {
	petP := &pet.Pet{}
	var petResponses []pet.Response

	if len(owner.Pets) > 0 {
		petResponses = petP.ToPetResponses(owner.Pets)
	}

	return &Response{
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

func (o *Owner) ToOwnerResponses(owners []Owner) []Response {
	ownerResponses := make([]Response, len(owners))
	for i, v := range owners {
		ownerResponses[i] = *o.ToOwnerResponse(&v)
	}
	return ownerResponses
}
