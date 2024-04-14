package vet

import (
	"encoding/json"
	"github.com/rhtran/golang-petclinic-service/pkg/infra/repository"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"

	"github.com/qiangxue/go-restful-api/pkg/log"
	"github.com/stretchr/testify/mock"
)

// vetServiceMock
type vetServiceMock struct {
	mock.Mock
}

func (vetM *vetServiceMock) GetVetById(id int) (*Response, error) {
	args := vetM.Called(id)
	intf := args.Get(0)
	val := intf.(Response)
	ptr := &val

	return ptr, args.Error(1)
}

func (vetM *vetServiceMock) GetVetByLastName(lastName string) ([]Response, error) {
	args := vetM.Called(lastName)
	intf := args.Get(0)
	val := intf.([]Response)
	ptr := val

	return ptr, args.Error(1)
}

func (vetM *vetServiceMock) GetAllVets() ([]Response, error) {
	args := vetM.Called()
	intf := args.Get(0)
	val := intf.([]Response)
	ptr := val

	return ptr, args.Error(1)
}

func (vetM *vetServiceMock) GetAllVetsWithSpecialties() ([]Response, error) {
	args := vetM.Called()
	intf := args.Get(0)
	val := intf.([]Response)
	ptr := val

	return ptr, args.Error(1)
}

func (vetM *vetServiceMock) Create(vet *repository.Vet) (*Response, error) {
	args := vetM.Called(vet)
	intf := args.Get(0)
	val := intf.(Response)
	ptr := &val

	return ptr, args.Error(1)
}

func (vetM *vetServiceMock) Update(vet *repository.Vet) (*Response, error) {
	args := vetM.Called(vet)
	intf := args.Get(0)
	val := intf.(Response)
	ptr := &val

	return ptr, args.Error(1)
}

// config the gin engine for testing purpose
func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	return r
}

func Test_VetById(t *testing.T) {
	logger := log.New().With(nil, "function", "Test_VetById")
	vetMock := vetServiceMock{}

	vetResponse := &Response{
		ID:        1,
		FirstName: "Nat",
		LastName:  "Cole",
	}

	vetMock.On("GetVetById", 1).Return(*vetResponse, nil)
	vetRouter := NewVetRouter(logger, &vetMock)

	r := setupRouter()
	v1 := r.Group("/v1")
	vetRouter.VetRegister(v1.Group("/vets"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/vets/1", nil)
	r.ServeHTTP(w, req)

	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)

	// unmarshal to Vet struct for asserts.
	actualVet := &repository.Vet{}
	json.Unmarshal(w.Body.Bytes(), actualVet)
	assert.Equal(t, vetResponse.ID, actualVet.ID)
	assert.Equal(t, vetResponse.FirstName, actualVet.FirstName)
	assert.Equal(t, vetResponse.LastName, actualVet.LastName)
}

func Test_VetByLastName(t *testing.T) {
	logger := log.New().With(nil, "function", "Test_VetByLastName")
	vetMock := vetServiceMock{}

	var vetResponses = make([]Response, 1)
	var vetResponse = &Response{
		ID:        15,
		FirstName: "Charles",
		LastName:  "Ward",
	}
	vetResponses[0] = *vetResponse

	vetMock.On("GetVetByLastName", "Ward").Return(vetResponses, nil)
	vetRouter := NewVetRouter(logger, &vetMock)

	r := setupRouter()
	v1 := r.Group("/v1")
	vetRouter.VetRegister(v1.Group("/vets"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/vets?last-name=Ward", nil)
	r.ServeHTTP(w, req)

	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)

	// unmarshal to Vet struct for asserts.
	actualVetResponses := make([]Response, 1)
	json.Unmarshal(w.Body.Bytes(), &actualVetResponses)

	assert.Equal(t, vetResponse.ID, actualVetResponses[0].ID)
	assert.Equal(t, vetResponse.FirstName, actualVetResponses[0].FirstName)
	assert.Equal(t, vetResponse.LastName, actualVetResponses[0].LastName)
}
