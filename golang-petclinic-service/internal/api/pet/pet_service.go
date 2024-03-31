package pet

import (
	"github.com/qiangxue/go-restful-api/pkg/log"
)

type Servicer interface {
	GetPetById(id int) (*Response, error)
	GetPetByName(name string) ([]Response, error)
	Create(pet *Pet) (*Response, error)
	Update(pet *Pet) (*Response, error)
}

type Service struct {
	logger     log.Logger
	repository Repositorier
}

func NewPetService(logger log.Logger, repository Repositorier) *Service {
	return &Service{logger: logger, repository: repository}
}

func (service *Service) GetPetById(id int) (*Response, error) {
	service.logger.Info("retrieve pet by id: %v", id)
	pet, err := service.repository.FindById(id)
	if err != nil {
		service.logger.Error("fails to retrieve pet by id: %v, err: %v", id, err.Error)
		return nil, err
	}

	petP := &Response{}
	petP.FromPet(pet)
	return petP, err
}

func (service *Service) GetPetByName(name string) ([]Response, error) {
	service.logger.Info("retrieve pet by name: %v", name)

	pets, err := service.repository.FindByName(name)
	if err != nil {
		service.logger.Error("fails to retrieve pet by name: %v, err: %v", name, err.Error)
		return nil, err
	}

	petP := &Pet{}
	return petP.ToPetResponses(pets), err
}

func (service *Service) Create(pet *Pet) (*Response, error) {
	service.logger.Info("Create new pet: %v", pet.Name)
	newPet, err := service.repository.Insert(pet)

	if err != nil {
		service.logger.Error("insert new pet failed: %v", err.Error())
		return &Response{}, err
	}

	petP := &Pet{}
	return petP.ToPetResponse(newPet), err
}

func (service *Service) Update(pet *Pet) (*Response, error) {
	service.logger.Info("update vet: %v", pet)
	updatedPet, err := service.repository.Update(pet)

	if err != nil {
		service.logger.Error("Update pet failed: %v", err.Error())
		return nil, err
	}

	petP := &Pet{}
	return petP.ToPetResponse(updatedPet), err
}
