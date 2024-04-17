package pet

import (
	"github.com/rhtran/golang-petclinic-service/pkg/infra/repository"
	"gorm.io/gorm"
	"time"
)

type Request struct {
	Name     string `json:"name" binding:"required"`
	Birthday string `json:"birthday" binding:"required"`
	TypeID   uint   `json:"typeId" binding:"required"`
	TypeName string `json:"typeName" binding:"required"`
}

func ToPet(petRequest *Request) *repository.Pet {
	birthday, _ := time.Parse(time.DateOnly, petRequest.Birthday)
	return &repository.Pet{
		Name:      petRequest.Name,
		BirthDate: &birthday,
		Type: repository.Type{
			Model: gorm.Model{
				ID: petRequest.TypeID,
			},
			Name: petRequest.TypeName,
		},
	}
}
