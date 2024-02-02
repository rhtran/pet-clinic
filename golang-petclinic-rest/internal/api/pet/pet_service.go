package pet

import (
	"github.com/qiangxue/go-restful-api/pkg/log"
)

type petRepository interface {
	FindById(id int) (*Pet, error)
	FindByName(name string) ([]Pet, error)
	Insert(pet *Pet) (*Pet, error)
	Update(pet *Pet) (*Pet, error)
}

type PetService struct {
	logger     log.Logger
	repository petRepository
}

func NewPetService(logger log.Logger, repository petRepository) *PetService {
	return &PetService{logger: logger, repository: repository}
}

func (service *PetService) GetPetById(id int) (*PetResponse, error) {
	service.logger.Info("retrieve pet by id: %v", id)
	pet, err := service.repository.FindById(id)
	if err != nil {
		service.logger.Error("fails to retrieve pet by id: %v, err: %v", id, err.Error)
		return nil, err
	}

	petP := &PetResponse{}
	petP.FromPet(pet)
	return petP, err
}

func (service *PetService) GetPetByName(name string) ([]PetResponse, error) {
	service.logger.Info("retrieve pet by name: %v", name)

	pets, err := service.repository.FindByName(name)
	if err != nil {
		service.logger.Error("fails to retrieve pet by name: %v, err: %v", name, err.Error)
		return nil, err
	}

	petP := &Pet{}
	return petP.ToPetResponses(pets), err
}

func (service *PetService) Create(pet *Pet) (*PetResponse, error) {
	service.logger.Info("Create new pet: %v", pet.Name)
	newPet, err := service.repository.Insert(pet)

	if err != nil {
		service.logger.Error("insert new pet failed: %v", err.Error())
		return &PetResponse{}, err
	}

	petP := &Pet{}
	return petP.ToPetResponse(newPet), err
}

func (service *PetService) Update(pet *Pet) (*PetResponse, error) {
	service.logger.Info("update vet: %v", pet)
	updatedPet, err := service.repository.Update(pet)

	if err != nil {
		service.logger.Error("Update pet failed: %v", err.Error())
		return nil, err
	}

	petP := &Pet{}
	return petP.ToPetResponse(updatedPet), err
}
