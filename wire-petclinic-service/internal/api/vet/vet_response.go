package vet

type VetResponse struct {
	ID          int         `json:"id"`
	FirstName   string      `json:"firstName"`
	LastName    string      `json:"lastName"`
	Specialties []Specialty `json:"specialties,omitempty"`
}
