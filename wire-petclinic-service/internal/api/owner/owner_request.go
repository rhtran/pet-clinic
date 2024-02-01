package owner

import (
	"wire-petclinic-service/pkg/model"
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
		Base: model.Base{
			ID: ownerRequest.ID,
		},
		Person: model.Person{
			FirstName: ownerRequest.FirstName,
			LastName:  ownerRequest.LastName,
		},

		Username:  ownerRequest.Username,
		Address:   ownerRequest.Address,
		City:      ownerRequest.City,
		Telephone: ownerRequest.Telephone,
	}
}
