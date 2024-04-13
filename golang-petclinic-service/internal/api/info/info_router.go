package info

import (
	"github.com/gin-gonic/gin"
	"github.com/qiangxue/go-restful-api/pkg/log"
	"net/http"
	"strings"
)

type Router struct {
	logger      log.Logger
	infoService Servicer
	ipService   IPServicer
}

func NewInfoRouter(logger log.Logger, infoService Servicer, ipService IPServicer) *Router {
	return &Router{logger,
		infoService, ipService}
}

// InfoRegister
// @Summary Show app info
// @Tags Info
// @Produce  json
// @Router /info [get]
// @Success 200 {object} info.Info
func (infoRouter *Router) InfoRegister(router *gin.RouterGroup) {
	router.GET("", infoRouter.appInfo)
}

// appInfo godoc
func (infoRouter *Router) appInfo(c *gin.Context) {
	result, err := infoRouter.infoService.getAppInfo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	addrs, err := infoRouter.ipService.lookupIP("localhost")
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
