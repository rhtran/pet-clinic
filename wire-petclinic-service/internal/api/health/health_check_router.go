package health

import (
	"github.com/qiangxue/go-restful-api/pkg/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckServicer interface {
	Check() (*HealthCheck, error)
}

type HealthCheckRouter struct {
	healthCheckService HealthCheckServicer
	logger             log.Logger
}

func NewHealthCheckRouter(healthCheckService HealthCheckServicer, logger log.Logger) *HealthCheckRouter {
	return &HealthCheckRouter{healthCheckService, logger}
}

func (healthCheckRouter *HealthCheckRouter) HealthCheckRegister(router *gin.RouterGroup) {
	router.GET("", healthCheckRouter.healthCheck)
}

func (healthCheckRouter *HealthCheckRouter) healthCheck(c *gin.Context) {
	result, err := healthCheckRouter.healthCheckService.Check()

	if err != nil {
		c.JSON(http.StatusBadRequest, result)
	} else {
		c.JSON(http.StatusOK, result)
	}
}
