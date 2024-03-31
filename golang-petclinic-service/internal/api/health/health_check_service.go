package health

import "github.com/qiangxue/go-restful-api/pkg/log"

type HealthCheckServicer interface {
	Check() (*Check, error)
}

type HealthCheckService struct {
	logger log.Logger
}

func NewHealthCheckService(logger log.Logger) *HealthCheckService {
	return &HealthCheckService{logger}
}

func (service *HealthCheckService) Check() (*Check, error) {
	healthCheck := &Check{
		Status: "UP",
	}

	return healthCheck, nil
}
