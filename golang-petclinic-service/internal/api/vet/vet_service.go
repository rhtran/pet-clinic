package vet

import (
	"github.com/qiangxue/go-restful-api/pkg/log"
	"github.com/rhtran/golang-petclinic-service/pkg/infra/repository/vet"
)

type VetService struct {
	logger     log.Logger
	repository vet.Repositorier
}

func NewVetService(logger log.Logger, repository vet.Repositorier) *VetService {
	return &VetService{logger: logger, repository: repository}
}

func (service *VetService) GetVetById(id int) (*Response, error) {
	getvet, err := service.repository.FindById(id)
	if err != nil {
		service.logger.Errorf("fails to retrieve vet by id: %d", id)
		return nil, err
	}

	response := &Response{}
	response.FromVet(getvet)
	return response, nil
}

func (service *VetService) GetVetByLastName(lastName string) ([]Response, error) {
	vets, err := service.repository.FindByLastName(lastName)
	if err != nil {
		service.logger.Errorf("fail to retrieve the vets by last name: %v, errors: %v", lastName, err.Error())
		return nil, err
	}

	return FromVets(vets), nil
}

func (service *VetService) GetAllVets() ([]Response, error) {
	vets, err := service.repository.FindAll()

	if err != nil {
		service.logger.Errorf("fail to retrieve all vets, errors: %v", err.Error())
		return nil, err
	}
	service.logger.Infof("counts of all vets: %d", len(vets))
	return FromVets(vets), nil
}

func (service *VetService) GetAllVetsWithSpecialties() ([]Response, error) {
	vets, err := service.repository.FindAllPreload()

	if err != nil {
		service.logger.Errorf("fail to retrieve all vets, errors: %v", err.Error())
		return nil, err
	}
	service.logger.Infof("counts of all vets: %d", len(vets))
	return FromVets(vets), nil
}

func (service *VetService) Create(vet *vet.Vet) (*Response, error) {
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

func (service *VetService) Update(vet *vet.Vet) (*Response, error) {
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
