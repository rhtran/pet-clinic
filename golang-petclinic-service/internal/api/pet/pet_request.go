package pet

import (
	"github.com/rhtran/golang-petclinic-service/pkg/infra/repository/pet"
	"github.com/rhtran/golang-petclinic-service/pkg/model"
)

type Request struct {
	Name     string `json:"name" binding:"required"`
	TypeID   int    `json:"typeId" binding:"required"`
	TypeName string `json:"typeName" binding:"required"`
}

func (pr Request) ToPet(petRequest *Request) *pet.Pet {
	return &pet.Pet{
		Name: petRequest.Name,
		Type: pet.Type{
			Base: model.Base{
				ID: petRequest.TypeID,
			},
			Name: petRequest.TypeName,
		},
	}
}
