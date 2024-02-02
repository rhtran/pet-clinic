package owner

import (
	"errors"
	resterr "github.com/rhtran/golang-petclinic-rest/internal/api/errors"
	"net/http"
	"strconv"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/qiangxue/go-restful-api/pkg/log"
)

type OwnerRouter struct {
	logger  log.Logger
	service Servicer
}

func NewOwnerRouter(logger log.Logger, service Servicer) *OwnerRouter {
	return &OwnerRouter{logger, service}
}

/*
 * Owner Router configuration
 */
func (ownerRouter *OwnerRouter) OwnerRegister(router *gin.RouterGroup) {
	router.GET(":id", ownerRouter.ownerById)
	router.GET(":id/pet", ownerRouter.allOwnersWithPets)
	router.GET("", ownerRouter.ownerByLastName)
	router.POST("", ownerRouter.addNewOwner)
	router.PUT(":id", ownerRouter.updateOwner)
}

func (ownerRouter *OwnerRouter) ownerById(c *gin.Context) {
	pathID := c.Param("id")

	if pathID == "all" {
		ownerRouter.allOwners(c)
		return
	}

	id, err := strconv.Atoi(pathID)
	if err != nil {
		c.JSON(http.StatusBadRequest, resterr.BadRequestWithDetails(err))
		return
	}

	response, err := ownerRouter.service.GetOwnerById(id)
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

func (ownerRouter *OwnerRouter) ownerByLastName(c *gin.Context) {
	lastName := c.Query("last-name")
	response, err := ownerRouter.service.GetOwnerByLastName(lastName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resterr.InternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response)
}

/*
Return all owners
*/
func (ownerRouter *OwnerRouter) allOwners(c *gin.Context) {
	responses, err := ownerRouter.service.GetAllOwners()
	if err != nil {
		c.JSON(http.StatusInternalServerError, resterr.InternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, responses)
}

func (ownerRouter *OwnerRouter) allOwnersWithPets(c *gin.Context) {
	responses, err := ownerRouter.service.GetAllOwnersWithPets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, resterr.InternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, responses)
}

func (ownerRouter *OwnerRouter) addNewOwner(c *gin.Context) {
	var ownerRequest OwnerRequest
	err := c.ShouldBindJSON(&ownerRequest)
	if err != nil {
		ownerRouter.logger.Errorf("Unable to Unmarshal JSON: %s", err.Error())
		c.JSON(http.StatusBadRequest, resterr.BadRequestWithDetails(err))
		return
	}

	newOwner, err := ownerRouter.service.Create(ownerRequest.ToOwner(&ownerRequest))
	c.JSON(http.StatusCreated, newOwner)
}

func (ownerRouter *OwnerRouter) updateOwner(c *gin.Context) {
	var ownerRequest OwnerRequest
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, resterr.BadRequestWithDetails(err))
		return
	}

	err = c.ShouldBindJSON(&ownerRequest)
	if err != nil {
		ownerRouter.logger.Errorf("Unable to Unmarshal JSON: %s", err.Error())
		c.JSON(http.StatusBadRequest, resterr.BadRequestWithDetails(err))
		return
	}

	ownerRequest.ID = id
	updatedOwner, err := ownerRouter.service.Update(ownerRequest.ToOwner(&ownerRequest))
	c.JSON(http.StatusCreated, updatedOwner)
}
