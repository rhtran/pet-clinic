package owner

import (
	model2 "github.com/rhtran/golang-petclinic-rest/pkg/model"
)

type OwnerRequest struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Address   string `json:"address"`
	City      string `json:"city"`
	Telephone string `json:"telephone"`
}

func (or *OwnerRequest) ToOwner(ownerRequest *OwnerRequest) *Owner {
	return &Owner{
		Base: model2.Base{
			ID: ownerRequest.ID,
		},
		Person: model2.Person{
			FirstName: ownerRequest.FirstName,
			LastName:  ownerRequest.LastName,
		},

		Username:  ownerRequest.Username,
		Address:   ownerRequest.Address,
		City:      ownerRequest.City,
		Telephone: ownerRequest.Telephone,
	}
}
