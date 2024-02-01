package vet

import (
	"github.com/qiangxue/go-restful-api/pkg/log"
)

type VetRepositorier interface {
	FindById(id int) (*Vet, error)
	FindByLastName(lastName string) ([]Vet, error)
	FindAll() ([]Vet, error)
	FindAllPreload() ([]Vet, error)
	Insert(vet *Vet) (*Vet, error)
	Update(vet *Vet) (*Vet, error)
}

type VetService struct {
	logger     log.Logger
	repository VetRepositorier
}

func NewVetService(logger log.Logger, repository VetRepositorier) *VetService {
	return &VetService{logger: logger, repository: repository}
}

func (service *VetService) GetVetById(id int) (*VetResponse, error) {
	vet, err := service.repository.FindById(id)
	if err != nil {
		service.logger.Errorf("fails to retrieve vet by id: %d", id)
		return nil, err
	}

	return vet.ToVetResponse(vet), nil
}

func (service *VetService) GetVetByLastName(lastName string) ([]VetResponse, error) {
	vets, err := service.repository.FindByLastName(lastName)
	if err != nil {
		service.logger.Errorf("fail to retrieve the vets by last name: %v, errors: %v", lastName, err.Error())
		return nil, err
	}

	vetP := &Vet{}
	return vetP.ToVetResponses(vets), nil
}

func (service *VetService) GetAllVets() ([]VetResponse, error) {
	vets, err := service.repository.FindAll()

	if err != nil {
		service.logger.Errorf("fail to retrieve all vets, errors: %v", err.Error())
		return nil, err
	}
	service.logger.Infof("counts of all vets: %d", len(vets))
	vetP := &Vet{}
	return vetP.ToVetResponses(vets), nil
}

func (service *VetService) GetAllVetsWithSpecialties() ([]VetResponse, error) {
	vets, err := service.repository.FindAllPreload()

	if err != nil {
		service.logger.Errorf("fail to retrieve all vets, errors: %v", err.Error())
		return nil, err
	}
	service.logger.Infof("counts of all vets: %d", len(vets))
	vetP := &Vet{}
	return vetP.ToVetResponses(vets), nil
}

func (service *VetService) Create(vet *Vet) (*VetResponse, error) {
	service.logger.Infof("Create new vet: %v", vet)
	newVet, err := service.repository.Insert(vet)

	if err != nil {
		service.logger.Errorf("insert new vet failed: %v", err.Error())
		return &VetResponse{}, err
	}

	vetP := &Vet{}
	return vetP.ToVetResponse(newVet), err
}

func (service *VetService) Update(vet *Vet) (*VetResponse, error) {
	service.logger.Infof("update vet: %v", vet)
	updatedVet, err := service.repository.Update(vet)

	if err != nil {
		service.logger.Errorf("Update vet failed: %v", err.Error())
		return nil, err
	}

	vetP := &Vet{}
	return vetP.ToVetResponse(updatedVet), err
}
