package pet

import (
	"github.com/rhtran/golang-petclinic-service/pkg/model"
)

type Request struct {
	Name     string `json:"name" binding:"required"`
	TypeID   int    `json:"typeId" binding:"required"`
	TypeName string `json:"typeName" binding:"required"`
}

func (pr Request) ToPet(petRequest *Request) *Pet {
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
