package health

import "github.com/qiangxue/go-restful-api/pkg/log"

type CheckServicer interface {
	check() (*Check, error)
}

type CheckService struct {
	logger log.Logger
}

func NewHealthCheckService(logger log.Logger) *CheckService {
	return &CheckService{logger}
}

func (service *CheckService) check() (*Check, error) {
	healthCheck := &Check{
		Status: "UP",
	}

	return healthCheck, nil
}
