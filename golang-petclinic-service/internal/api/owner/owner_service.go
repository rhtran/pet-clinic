package owner

import (
	"github.com/qiangxue/go-restful-api/pkg/log"
)

type Servicer interface {
	GetOwnerById(id int) (*Response, error)
	GetOwnerByLastName(lastName string) ([]Response, error)
	GetAllOwners() ([]Response, error)
	GetAllOwnersWithPets() ([]Response, error)
	Create(owner *Owner) (*Response, error)
	Update(owner *Owner) (*Response, error)
}

type Service struct {
	logger     log.Logger
	repository Repositorier
}

func NewOwnerService(logger log.Logger, repository Repositorier) *Service {
	return &Service{logger: logger, repository: repository}
}

func (service *Service) GetOwnerById(id int) (*Response, error) {
	service.logger.Infof("retrieve owner by id: %v", id)
	owner, err := service.repository.FindById(id)
	if err != nil {
		service.logger.Errorf("fails to retrieve owner by id: %v, err: %v", id, err.Error)
		return nil, err
	}

	ownerP := &Owner{}
	return ownerP.ToOwnerResponse(owner), err
}

func (service *Service) GetOwnerByLastName(lastName string) ([]Response, error) {
	service.logger.Infof("retrieve owners by last name: %v", lastName)
	owners, err := service.repository.FindByLastName(lastName)
	if err != nil {
		service.logger.Errorf("fails to retrieve owner by last name: %v, err: %v", lastName, err.Error)
		return nil, err
	}

	ownerP := &Owner{}
	return ownerP.ToOwnerResponses(owners), err
}

func (service *Service) GetAllOwners() ([]Response, error) {
	service.logger.Info("retrieve all owner")
	owners, err := service.repository.FindAll()
	if err != nil {
		service.logger.Errorf("fails to retrieve all owner, err: %v", err.Error)
		return nil, err
	}

	ownerP := &Owner{}
	return ownerP.ToOwnerResponses(owners), err
}

func (service *Service) GetAllOwnersWithPets() ([]Response, error) {
	service.logger.Info("retrieve all owner")
	owners, err := service.repository.FindAllWithPets()
	if err != nil {
		service.logger.Errorf("fails to retrieve all owner, err: %v", err.Error)
		return nil, err
	}

	ownerP := &Owner{}
	return ownerP.ToOwnerResponses(owners), err
}

func (service *Service) Create(owner *Owner) (*Response, error) {
	service.logger.Info("Create new owner")
	newOwner, err := service.repository.Insert(owner)

	if err != nil {
		service.logger.Errorf("insert new owner failed: %v", err.Error())
		return nil, err
	}

	ownerP := &Owner{}
	return ownerP.ToOwnerResponse(newOwner), err
}

func (service *Service) Update(owner *Owner) (*Response, error) {
	service.logger.Infof("update vet: %v", owner)
	updatedOwner, err := service.repository.Update(owner)

	if err != nil {
		service.logger.Errorf("Update owner failed: %v", err.Error())
		return nil, err
	}

	ownerP := &Owner{}
	return ownerP.ToOwnerResponse(updatedOwner), err
}
