package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qiangxue/go-restful-api/pkg/log"
)

type CheckRouter struct {
	healthCheckService CheckServicer
	logger             log.Logger
}

func NewHealthCheckRouter(healthCheckService CheckServicer, logger log.Logger) *CheckRouter {
	return &CheckRouter{healthCheckService, logger}
}

func (healthCheckRouter *CheckRouter) HealthCheckRegister(router *gin.RouterGroup) {
	router.GET("", healthCheckRouter.healthCheck)
}

func (healthCheckRouter *CheckRouter) healthCheck(c *gin.Context) {
	result, err := healthCheckRouter.healthCheckService.check()

	if err != nil {
		c.JSON(http.StatusBadRequest, result)
	} else {
		c.JSON(http.StatusOK, result)
	}
}
