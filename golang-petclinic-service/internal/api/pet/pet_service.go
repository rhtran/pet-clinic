package pet

import (
	"github.com/qiangxue/go-restful-api/pkg/log"
	"github.com/rhtran/golang-petclinic-service/pkg/infra/repository/pet"
)

type Servicer interface {
	GetPetById(id int) (*Response, error)
	GetPetByName(name string) ([]Response, error)
	Create(pet *pet.Pet) (*Response, error)
	Update(pet *pet.Pet) (*Response, error)
}

type Service struct {
	logger     log.Logger
	repository pet.Repositorier
}

func NewPetService(logger log.Logger, repository pet.Repositorier) *Service {
	return &Service{logger: logger, repository: repository}
}

func (service *Service) GetPetById(id int) (*Response, error) {
	service.logger.Info("retrieve pet by id: %v", id)
	petF, err := service.repository.FindById(id)
	if err != nil {
		service.logger.Error("fails to retrieve pet by id: %v, err: %v", id, err.Error)
		return nil, err
	}

	response := &Response{}
	response.FromPet(petF)
	return response, nil
}

func (service *Service) GetPetByName(name string) ([]Response, error) {
	service.logger.Info("retrieve pet by name: %v", name)

	pets, err := service.repository.FindByName(name)
	if err != nil {
		service.logger.Error("fails to retrieve pet by name: %v, err: %v", name, err.Error)
		return nil, err
	}

	return FromPets(pets), nil
}

func (service *Service) Create(pet *pet.Pet) (*Response, error) {
	service.logger.Info("Create new pet: %v", pet.Name)
	newPet, err := service.repository.Insert(pet)

	if err != nil {
		service.logger.Error("insert new pet failed: %v", err.Error())
		return &Response{}, err
	}

	response := &Response{}
	response.FromPet(newPet)
	return response, nil
}

func (service *Service) Update(pet *pet.Pet) (*Response, error) {
	service.logger.Info("update vet: %v", pet)
	updatedPet, err := service.repository.Update(pet)

	if err != nil {
		service.logger.Error("Update pet failed: %v", err.Error())
		return nil, err
	}

	response := &Response{}
	response.FromPet(updatedPet)
	return response, nil
}
