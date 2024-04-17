package owner

import (
	"github.com/rhtran/golang-petclinic-service/pkg/infra/repository"
	"github.com/rhtran/golang-petclinic-service/pkg/model"
)

type Request struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Address   string `json:"address"`
	City      string `json:"city"`
	Telephone string `json:"telephone"`
}

func ToOwner(ownerRequest *Request) *repository.Owner {
	return &repository.Owner{
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
