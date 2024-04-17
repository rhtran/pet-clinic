package pet

import (
	"github.com/rhtran/golang-petclinic-service/pkg/infra/repository"
	"time"
)

type Response struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	BirthDate string `json:"birthDate"`
	Type      string `json:"type"`
}

func (pr *Response) FromPet(pet *repository.Pet) {
	pr.ID = pet.ID
	pr.Name = pet.Name
	pr.BirthDate = pet.BirthDate.Format(time.DateOnly)
	pr.Type = pet.Type.Name
}

func FromPets(pets []repository.Pet) []Response {
	petResponses := make([]Response, len(pets))
	for i, v := range pets {
		petResponses[i].FromPet(&v)
	}
	return petResponses
}
