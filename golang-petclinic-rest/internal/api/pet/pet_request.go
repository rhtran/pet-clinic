package pet

import (
	"github.com/rhtran/golang-petclinic-rest/pkg/model"
)

type PetRequest struct {
	Name     string `json:"name" binding:"required"`
	TypeID   int    `json:"typeId" binding:"required"`
	TypeName string `json:"typeName" binding:"required"`
}

func (pr PetRequest) ToPet(petRequest *PetRequest) *Pet {
	return &Pet{
		Name: petRequest.Name,
		Type: Type{
			Base: model.Base{
				ID: petRequest.TypeID,
			},
			Name: petRequest.TypeName,
		},
	}
}
