package owner

import (
	"encoding/json"
	"github.com/rhtran/golang-petclinic-service/internal/api/test"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/qiangxue/go-restful-api/pkg/log"
)

// config the gin engine for testing purpose
func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	return r
}

func Test_ownerById(t *testing.T) {
	logger := log.New().With(nil, "function", "Test_ownerById")
	ownerMock := MockServicer{}

	ownerResponse := &Response{
		ID:        1,
		FirstName: "Nat",
		LastName:  "Cole",
	}

	ownerMock.On("getOwnerById", 1).Return(ownerResponse, nil)
	ownerRouter := NewOwnerRouter(logger, &ownerMock)

	r := setupRouter()
	v1 := r.Group("/v1")
	ownerRouter.OwnerRegister(v1.Group("/owners"))

	tc1, _ := json.Marshal(ownerResponse)

	tests := []test.APITestCase{
		{"Get Owner By ID", "GET", "/v1/owners/1", "", nil, http.StatusOK, string(tc1)},
	}
	for _, tc := range tests {
		test.Endpoint(t, r, tc)
	}
}

func Test_OwnerByLastName(t *testing.T) {
	logger := log.New().With(nil, "function", "Test_OwnerByLastName")
	ownerMock := MockServicer{}

	var ownerResponses = make([]Response, 1)
	ownerResponse := &Response{
		ID:        15,
		FirstName: "Charles",
		LastName:  "Ward",
	}
	ownerResponses[0] = *ownerResponse

	ownerMock.On("getOwnerByLastName", "Ward").Return(ownerResponses, nil)
	ownerRouter := NewOwnerRouter(logger, &ownerMock)

	r := setupRouter()
	v1 := r.Group("/v1")
	ownerRouter.OwnerRegister(v1.Group("/owners"))

	tc1, _ := json.Marshal(ownerResponses)

	tests := []test.APITestCase{
		{"Get Owner By Last Name", "GET", "/v1/owners?last-name=Ward", "", nil, http.StatusOK, string(tc1)},
	}
	for _, tc := range tests {
		test.Endpoint(t, r, tc)
	}
}
