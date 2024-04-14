package vet

import (
	"github.com/qiangxue/go-restful-api/pkg/log"
	"github.com/rhtran/golang-petclinic-service/pkg/infra/repository"
)

type Servicer interface {
	GetVetById(id int) (*Response, error)
	GetVetByLastName(lastName string) ([]Response, error)
	GetAllVets() ([]Response, error)
	GetAllVetsWithSpecialties() ([]Response, error)
	Create(vet *repository.Vet) (*Response, error)
	Update(vet *repository.Vet) (*Response, error)
}

type Service struct {
	logger     log.Logger
	repository repository.VetRepositorier
}

func NewVetService(logger log.Logger, repository repository.VetRepositorier) *Service {
	return &Service{logger: logger, repository: repository}
}

func (service *Service) GetVetById(id int) (*Response, error) {
	getvet, err := service.repository.FindById(id)
	if err != nil {
		service.logger.Errorf("fails to retrieve vet by id: %d", id)
		return nil, err
	}

	response := &Response{}
	response.FromVet(getvet)
	return response, nil
}

func (service *Service) GetVetByLastName(lastName string) ([]Response, error) {
	vets, err := service.repository.FindByLastName(lastName)
	if err != nil {
		service.logger.Errorf("fail to retrieve the vets by last name: %v, errors: %v", lastName, err.Error())
		return nil, err
	}

	return FromVets(vets), nil
}

func (service *Service) GetAllVets() ([]Response, error) {
	vets, err := service.repository.FindAll()

	if err != nil {
		service.logger.Errorf("fail to retrieve all vets, errors: %v", err.Error())
		return nil, err
	}
	service.logger.Infof("counts of all vets: %d", len(vets))
	return FromVets(vets), nil
}

func (service *Service) GetAllVetsWithSpecialties() ([]Response, error) {
	vets, err := service.repository.FindAllPreload()

	if err != nil {
		service.logger.Errorf("fail to retrieve all vets, errors: %v", err.Error())
		return nil, err
	}
	service.logger.Infof("counts of all vets: %d", len(vets))
	return FromVets(vets), nil
}

func (service *Service) Create(vet *repository.Vet) (*Response, error) {
	service.logger.Infof("Create new vet: %v", vet)
	newVet, err := service.repository.Insert(vet)

	if err != nil {
		service.logger.Errorf("insert new vet failed: %v", err.Error())
		return &Response{}, err
	}

	response := &Response{}
	response.FromVet(newVet)
	return response, nil
}

func (service *Service) Update(vet *repository.Vet) (*Response, error) {
	service.logger.Infof("update vet: %v", vet)
	updatedVet, err := service.repository.Update(vet)

	if err != nil {
		service.logger.Errorf("Update vet failed: %v", err.Error())
		return nil, err
	}

	response := &Response{}
	response.FromVet(updatedVet)
	return response, nil
}
