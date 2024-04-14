package owner

import (
	"errors"
	resterr "github.com/rhtran/golang-petclinic-service/middleware/errors"
	"net/http"
	"strconv"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/qiangxue/go-restful-api/pkg/log"
)

type Router struct {
	logger  log.Logger
	service Servicer
}

func NewOwnerRouter(logger log.Logger, service Servicer) *Router {
	return &Router{logger, service}
}

/*
 * Owner Router configuration
 */
func (ownerRouter *Router) OwnerRegister(router *gin.RouterGroup) {
	router.GET(":id", ownerRouter.ownerById)
	router.GET(":id/pet", ownerRouter.allOwnersWithPets)
	router.GET("", ownerRouter.ownerByLastName)
	router.POST("", ownerRouter.addNewOwner)
	router.PUT(":id", ownerRouter.updateOwner)
}

func (ownerRouter *Router) ownerById(c *gin.Context) {
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

	response, err := ownerRouter.service.getOwnerById(id)
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

func (ownerRouter *Router) ownerByLastName(c *gin.Context) {
	lastName := c.Query("last-name")
	response, err := ownerRouter.service.getOwnerByLastName(lastName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resterr.InternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response)
}

/*
Return all owners
*/
func (ownerRouter *Router) allOwners(c *gin.Context) {
	responses, err := ownerRouter.service.getAllOwners()
	if err != nil {
		c.JSON(http.StatusInternalServerError, resterr.InternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, responses)
}

func (ownerRouter *Router) allOwnersWithPets(c *gin.Context) {
	responses, err := ownerRouter.service.getAllOwnersWithPets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, resterr.InternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, responses)
}

func (ownerRouter *Router) addNewOwner(c *gin.Context) {
	var ownerRequest Request
	err := c.ShouldBindJSON(&ownerRequest)
	if err != nil {
		ownerRouter.logger.Errorf("Unable to Unmarshal JSON: %s", err.Error())
		c.JSON(http.StatusBadRequest, resterr.BadRequestWithDetails(err))
		return
	}

	newOwner, err := ownerRouter.service.create(ToOwner(&ownerRequest))
	c.JSON(http.StatusCreated, newOwner)
}

func (ownerRouter *Router) updateOwner(c *gin.Context) {
	var ownerRequest Request
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
	updatedOwner, err := ownerRouter.service.update(ToOwner(&ownerRequest))
	c.JSON(http.StatusCreated, updatedOwner)
}
