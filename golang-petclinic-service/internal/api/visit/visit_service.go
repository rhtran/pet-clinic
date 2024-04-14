package visit

import (
	"github.com/qiangxue/go-restful-api/pkg/log"
	"github.com/rhtran/golang-petclinic-service/pkg/infra/repository"
)

type Servicer interface {
	getVisitById(id int) (*Response, error)
	getAllVisits() ([]Response, error)
}

type Service struct {
	logger     log.Logger
	repository repository.VisitRepositorier
}

func NewVisitService(logger log.Logger, repository repository.VisitRepositorier) *Service {
	return &Service{logger: logger, repository: repository}
}

func (service *Service) getVisitById(id int) (*Response, error) {
	visit, err := service.repository.FindById(id)
	if err != nil {
		service.logger.Errorf("fails to retrieve visit by id: %d", id)
		return nil, err
	}

	response := &Response{}
	response.FromVisit(visit)
	return response, nil
}

func (service *Service) getAllVisits() ([]Response, error) {
	visits, err := service.repository.FindAll()

	if err != nil {
		service.logger.Errorf("fail to retrieve all visits, errors: %v", err.Error())
		return nil, err
	}
	service.logger.Infof("counts of all visits: %d", len(visits))
	return FromVisits(visits), nil
}
