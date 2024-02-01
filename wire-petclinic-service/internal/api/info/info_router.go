package info

import (
	"net"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/qiangxue/go-restful-api/pkg/log"
)

type ipLookupService interface {
	LookupIP(host string) ([]net.IP, error)
}

type Router struct {
	logger          log.Logger
	infoService     Servicer
	ipLookupService ipLookupService
}

func NewInfoRouter(logger log.Logger, infoService Servicer, ipLookupService ipLookupService) *Router {
	return &Router{logger,
		infoService,
		ipLookupService}
}

// ShowInfo godoc
// @Summary Show app info
// @Tags Info
// @Produce  json
// @Router /info [get]
// @Success 200 {object} info.Info
func (infoRouter *Router) InfoRegister(router *gin.RouterGroup) {
	router.GET("", infoRouter.appInfo)
}

func (infoRouter *Router) appInfo(c *gin.Context) {
	result, err := infoRouter.infoService.GetAppInfo()

	// this line make this function mat testable.
	addrs, err := net.LookupIP("localhost")
	if err != nil {
		result.Ip = "Unknown host"
	} else {
		for _, ia := range addrs {
			result.Ip += ia.String() + " "
		}
		result.Ip = strings.TrimSpace(result.Ip)
	}

	c.JSON(http.StatusOK, result)
}
