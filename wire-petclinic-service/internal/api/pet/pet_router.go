package pet

import (
	"errors"
	"github.com/qiangxue/go-restful-api/pkg/log"
	"net/http"
	"strconv"
	"wire-petclinic-service/internal/api"
	resterr "wire-petclinic-service/internal/api/errors"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type petService interface {
	GetPetById(id int) (*PetResponse, error)
	GetPetByName(name string) ([]PetResponse, error)
	Create(pet *Pet) (*PetResponse, error)
	Update(pet *Pet) (*PetResponse, error)
}

type PetRouter struct {
	logger  log.Logger
	service petService
}

func NewPetRouter(logger log.Logger, service petService) *PetRouter {
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
	var request PetRequest
	err := c.ShouldBind(&request)
	if err != nil {
		petRouter.logger.Errorf("Unable to Unmarshal JSON: %s", err.Error())
		c.JSON(http.StatusBadRequest, resterr.BadRequestWithDetails(err))
		return
	}

	petRouter.logger.Infof("Add a new pet: %v", request.Name)
	petResponse, err := petRouter.service.Create(request.ToPet(&request))
	c.JSON(http.StatusCreated, petResponse)
}

func (petRouter *PetRouter) updatePet(c *gin.Context) {
	var request PetRequest
	err := c.ShouldBind(&request)
	if err != nil {
		petRouter.logger.Errorf("Unable to Unmarshal JSON: %s", err.Error())
		c.JSON(http.StatusBadRequest, resterr.BadRequestWithDetails(err))
		return
	}

	petRouter.logger.Infof("Add a new pet: %v", request.Name)
	petResponse, err := petRouter.service.Update(request.ToPet(&request))
	c.JSON(http.StatusCreated, petResponse)
}
