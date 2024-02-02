package pet

import "time"

type PetResponse struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	BirthDate *time.Time `json:"birthDate"`
	Type      string     `json:"type"`
}

func (pr *PetResponse) FromPet(pet *Pet) {
	pr.ID = pet.ID
	pr.Name = pet.Name
	pr.BirthDate = pet.BirthDate
	pr.Type = pet.Type.Name
}
