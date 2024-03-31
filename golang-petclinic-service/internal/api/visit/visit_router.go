package visit

import (
	"github.com/rhtran/golang-petclinic-service/middleware/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qiangxue/go-restful-api/pkg/log"
)

type VisitRouter struct {
	logger  log.Logger
	service Servicer
}

func NewVisitRouter(logger log.Logger, service Servicer) *VisitRouter {
	return &VisitRouter{logger, service}
}

func (visitRouter *VisitRouter) VisitRegister(router *gin.RouterGroup) {
	router.GET(":id", visitRouter.visitById)
}

func (visitRouter *VisitRouter) visitById(c *gin.Context) {
	pathID := c.Param("id")
	if pathID == "all" {
		visitRouter.allVisits(c)
		return
	}

	id, err := strconv.Atoi(pathID)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.BadRequest(err.Error()))
		return
	}

	response, err := visitRouter.service.getVisitById(id)
	if err != nil {

		c.JSON(http.StatusInternalServerError, errors.InternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response)
}

func (visitRouter *VisitRouter) allVisits(c *gin.Context) {
	response, err := visitRouter.service.getAllVisits()
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, response)
}
