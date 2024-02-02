package health

import "github.com/qiangxue/go-restful-api/pkg/log"

type HealthCheckService struct {
	logger log.Logger
}

func NewHealthCheckService(logger log.Logger) *HealthCheckService {
	return &HealthCheckService{logger}
}

func (service *HealthCheckService) Check() (*HealthCheck, error) {
	healthCheck := &HealthCheck{
		Status: "UP",
	}

	return healthCheck, nil
}
