package owner

import (
	"github.com/rhtran/golang-petclinic-service/infra/repository/owner"
	model2 "github.com/rhtran/golang-petclinic-service/pkg/model"
)

type Request struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Address   string `json:"address"`
	City      string `json:"city"`
	Telephone string `json:"telephone"`
}

func (or *Request) ToOwner(ownerRequest *Request) *owner.Owner {
	return &owner.Owner{
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
