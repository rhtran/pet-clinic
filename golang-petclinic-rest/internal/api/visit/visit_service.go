package visit

import "github.com/qiangxue/go-restful-api/pkg/log"

type VisitRepositorier interface {
	FindById(id int) (*Visit, error)
	FindAll() ([]Visit, error)
}

type VisitService struct {
	logger     log.Logger
	repository VisitRepositorier
}

func NewVisitService(logger log.Logger, repository VisitRepositorier) *VisitService {
	return &VisitService{logger: logger, repository: repository}
}

func (service *VisitService) GetVisitById(id int) (*VisitResponse, error) {
	visit, err := service.repository.FindById(id)
	if err != nil {
		service.logger.Errorf("fails to retrieve visit by id: %d", id)
		return nil, err
	}

	return visit.ToVisitResponse(visit), nil
}

func (service *VisitService) GetAllVisits() ([]VisitResponse, error) {
	visits, err := service.repository.FindAll()

	if err != nil {
		service.logger.Errorf("fail to retrieve all visits, errors: %v", err.Error())
		return nil, err
	}
	service.logger.Infof("counts of all visits: %d", len(visits))
	visitP := &Visit{}
	return visitP.ToVisitResponses(visits), nil
}
