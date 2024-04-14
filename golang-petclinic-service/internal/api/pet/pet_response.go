package pet

import (
	"github.com/rhtran/golang-petclinic-service/pkg/infra/repository/pet"
	"time"
)

type Response struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	BirthDate *time.Time `json:"birthDate"`
	Type      string     `json:"type"`
}

func (pr *Response) FromPet(pet *pet.Pet) {
	pr.ID = pet.ID
	pr.Name = pet.Name
	pr.BirthDate = pet.BirthDate
	pr.Type = pet.Type.Name
}

func FromPets(pets []pet.Pet) []Response {
	petResponses := make([]Response, len(pets))
	for i, v := range pets {
		petResponses[i].FromPet(&v)
	}
	return petResponses
}
