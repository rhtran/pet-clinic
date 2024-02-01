package visit

import (
	"net/http"
	"strconv"
	"wire-petclinic-service/internal/api/errors"

	"github.com/gin-gonic/gin"
	"github.com/qiangxue/go-restful-api/pkg/log"
)

type VisitServicer interface {
	GetVisitById(id int) (*VisitResponse, error)
	GetAllVisits() ([]VisitResponse, error)
}

type VisitRouter struct {
	logger  log.Logger
	service VisitServicer
}

func NewVisitRouter(logger log.Logger, service VisitServicer) *VisitRouter {
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

	response, err := visitRouter.service.GetVisitById(id)
	if err != nil {

		c.JSON(http.StatusInternalServerError, errors.InternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response)
}

func (visitRouter *VisitRouter) allVisits(c *gin.Context) {
	response, err := visitRouter.service.GetAllVisits()
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, response)
}
