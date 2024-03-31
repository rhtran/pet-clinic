package vet

type Response struct {
	ID          int         `json:"id"`
	FirstName   string      `json:"firstName"`
	LastName    string      `json:"lastName"`
	Specialties []Specialty `json:"specialties,omitempty"`
}

type SomeResponse struct {
	Vets []Response `json:"vets"`
}
