package vet

import (
	"errors"
	resterr "github.com/rhtran/golang-petclinic-service/middleware/errors"
	"net/http"
	"strconv"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/qiangxue/go-restful-api/pkg/log"
)

type VetRouter struct {
	logger  log.Logger
	service Servicer
}

func NewVetRouter(logger log.Logger, service Servicer) *VetRouter {
	return &VetRouter{logger, service}
}

func (vetRouter *VetRouter) VetRegister(router *gin.RouterGroup) {
	router.GET(":id", vetRouter.vetById)
	router.GET(":id/specialty", vetRouter.AllVetsWithSpecialties)
	router.GET("", vetRouter.vetByParam)
	router.POST("", vetRouter.addNewVet)
	router.PUT(":id", vetRouter.updateVet)
}

func (vetRouter *VetRouter) vetById(c *gin.Context) {
	pathID := c.Param("id")
	if pathID == "all" {
		vetRouter.AllVets(c)
		return
	}

	id, err := strconv.Atoi(pathID)
	if err != nil {
		c.JSON(http.StatusBadRequest, resterr.BadRequest(err.Error()))
		return
	}

	response, err := vetRouter.service.getVetById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, resterr.NotFound(err.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, resterr.InternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response)
}

func (vetRouter *VetRouter) vetByParam(c *gin.Context) {
	lastName := c.Query("last-name")

	vetRouter.vetByLastName(c, lastName)
	return
}

func (vetRouter *VetRouter) vetByLastName(c *gin.Context, lastName string) {
	response, err := vetRouter.service.getVetByLastName(lastName)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, response)
}

func (vetRouter *VetRouter) AllVets(c *gin.Context) {
	response, err := vetRouter.service.getAllVets()
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, response)
}

func (vetRouter *VetRouter) AllVetsWithSpecialties(c *gin.Context) {
	response, err := vetRouter.service.getAllVetsWithSpecialties()
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, response)
}

func (vetRouter *VetRouter) addNewVet(c *gin.Context) {
	var vetRequest Request
	err := c.ShouldBindJSON(&vetRequest)
	if err != nil {
		vetRouter.logger.Errorf("Unable to Unmarshal JSON: %s", err.Error())
		c.JSON(http.StatusBadRequest, resterr.BadRequestWithDetails(err))
		return
	}

	newVet, err := vetRouter.service.create(ToVet(&vetRequest))
	c.JSON(http.StatusCreated, newVet)
}

func (vetRouter *VetRouter) updateVet(c *gin.Context) {
	var vetRequest Request
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	err := c.ShouldBindJSON(&vetRequest)

	if err != nil {
		vetRouter.logger.Error("Unable to Unmarshal JSON: %s", err.Error())
		c.JSON(http.StatusBadRequest, resterr.BadRequestWithDetails(err))
		return
	}

	vetEntity := ToVet(&vetRequest)
	vetEntity.ID = uint(id)
	newVet, err := vetRouter.service.update(vetEntity)
	c.JSON(http.StatusCreated, newVet)
}
