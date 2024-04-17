package pet

import (
	"errors"
	"github.com/rhtran/golang-petclinic-service/internal/api"
	resterr "github.com/rhtran/golang-petclinic-service/middleware/errors"
	"net/http"
	"strconv"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/qiangxue/go-restful-api/pkg/log"
)

type PetRouter struct {
	logger  log.Logger
	service Servicer
}

func NewPetRouter(logger log.Logger, service Servicer) *PetRouter {
	return &PetRouter{logger, service}
}

func (petRouter *PetRouter) PetRegister(router *gin.RouterGroup) {
	router.GET(":id", petRouter.petById)
	router.GET("", petRouter.petByQueryParam)
	router.POST("", petRouter.addNewPet)
	router.PUT("", petRouter.updatePet)
}

func (petRouter *PetRouter) petById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, resterr.BadRequest(err.Error()))
		return
	}

	response, err := petRouter.service.GetPetById(id)
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

func (petRouter *PetRouter) petByQueryParam(c *gin.Context) {
	var nameParam api.NameParam

	err := c.BindQuery(&nameParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, resterr.BadRequestWithDetails(err))
		return
	}

	petRouter.petByName(c, nameParam)
}

func (petRouter *PetRouter) petByName(c *gin.Context, param api.NameParam) {
	response, err := petRouter.service.GetPetByName(param.Name)

	if err != nil {
		c.JSON(http.StatusInternalServerError, resterr.InternalServerError(""))
		return
	}

	c.JSON(http.StatusOK, response)
}

func (petRouter *PetRouter) addNewPet(c *gin.Context) {
	var request Request
	err := c.ShouldBind(&request)
	if err != nil {
		petRouter.logger.Errorf("Unable to Unmarshal JSON: %s", err.Error())
		c.JSON(http.StatusBadRequest, resterr.BadRequestWithDetails(err))
		return
	}

	petRouter.logger.Infof("Add a new pet: %v", request.Name)
	petResponse, err := petRouter.service.Create(ToPet(&request))
	c.JSON(http.StatusCreated, petResponse)
}

func (petRouter *PetRouter) updatePet(c *gin.Context) {
	var request Request
	err := c.ShouldBind(&request)
	if err != nil {
		petRouter.logger.Errorf("Unable to Unmarshal JSON: %s", err.Error())
		c.JSON(http.StatusBadRequest, resterr.BadRequestWithDetails(err))
		return
	}

	petRouter.logger.Infof("Add a new pet: %v", request.Name)
	petResponse, err := petRouter.service.Update(ToPet(&request))
	c.JSON(http.StatusCreated, petResponse)
}
