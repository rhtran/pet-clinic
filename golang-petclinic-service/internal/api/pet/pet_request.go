package pet

import (
	"github.com/rhtran/golang-petclinic-service/pkg/infra/repository"
	"gorm.io/gorm"
)

type Request struct {
	Name     string `json:"name" binding:"required"`
	TypeID   uint   `json:"typeId" binding:"required"`
	TypeName string `json:"typeName" binding:"required"`
}

func ToPet(petRequest *Request) *repository.Pet {
	return &repository.Pet{
		Name: petRequest.Name,
		Type: repository.Type{
			Model: gorm.Model{
				ID: petRequest.TypeID,
			},
			Name: petRequest.TypeName,
		},
	}
}
