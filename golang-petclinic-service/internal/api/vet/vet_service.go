package vet

import (
	"github.com/qiangxue/go-restful-api/pkg/log"
	"github.com/rhtran/golang-petclinic-service/pkg/infra/repository"
)

type Servicer interface {
	getVetById(id int) (*Response, error)
	getVetByLastName(lastName string) ([]Response, error)
	getAllVets() ([]Response, error)
	getAllVetsWithSpecialties() ([]Response, error)
	create(vet *repository.Vet) (*Response, error)
	update(vet *repository.Vet) (*Response, error)
}

type Service struct {
	logger     log.Logger
	repository repository.VetRepositorier
}

func NewVetService(logger log.Logger, repository repository.VetRepositorier) *Service {
	return &Service{logger: logger, repository: repository}
}

func (service *Service) getVetById(id int) (*Response, error) {
	getvet, err := service.repository.FindById(id)
	if err != nil {
		service.logger.Errorf("fails to retrieve vet by id: %d", id)
		return nil, err
	}

	response := &Response{}
	response.FromVet(getvet)
	return response, nil
}

func (service *Service) getVetByLastName(lastName string) ([]Response, error) {
	vets, err := service.repository.FindByLastName(lastName)
	if err != nil {
		service.logger.Errorf("fail to retrieve the vets by last name: %v, errors: %v", lastName, err.Error())
		return nil, err
	}

	return FromVets(vets), nil
}

func (service *Service) getAllVets() ([]Response, error) {
	vets, err := service.repository.FindAll()

	if err != nil {
		service.logger.Errorf("fail to retrieve all vets, errors: %v", err.Error())
		return nil, err
	}
	service.logger.Infof("counts of all vets: %d", len(vets))
	return FromVets(vets), nil
}

func (service *Service) getAllVetsWithSpecialties() ([]Response, error) {
	vets, err := service.repository.FindAllPreload()

	if err != nil {
		service.logger.Errorf("fail to retrieve all vets, errors: %v", err.Error())
		return nil, err
	}
	service.logger.Infof("counts of all vets: %d", len(vets))
	return FromVets(vets), nil
}

func (service *Service) create(vet *repository.Vet) (*Response, error) {
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

func (service *Service) update(vet *repository.Vet) (*Response, error) {
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
