package pet

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"

	"github.com/qiangxue/go-restful-api/pkg/log"
	"github.com/stretchr/testify/mock"
)

// petServiceMock
// require to mock all functions; otherwise, compilation errors.
type petServiceMock struct {
	mock.Mock
}

func (petM *petServiceMock) GetPetById(id int) (*PetResponse, error) {
	args := petM.Called(id)
	intf := args.Get(0)
	val := intf.(PetResponse)
	ptr := &val

	return ptr, args.Error(1)
}

func (petM *petServiceMock) GetPetByName(name string) ([]PetResponse, error) {
	args := petM.Called(name)
	intf := args.Get(0)
	val := intf.([]PetResponse)
	ptr := val

	return ptr, args.Error(1)
}

func (petM *petServiceMock) Create(pet *Pet) (*PetResponse, error) {
	args := petM.Called(pet)
	intf := args.Get(0)
	val := intf.(*PetResponse)
	ptr := val

	return ptr, args.Error(1)
}

func (petM *petServiceMock) Update(pet *Pet) (*PetResponse, error) {
	args := petM.Called(pet)
	intf := args.Get(0)
	val := intf.(*PetResponse)
	ptr := val

	return ptr, args.Error(1)
}

// config the gin engine for testing purpose
func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	return r
}

func Test_PetById(t *testing.T) {
	logger := log.New().With(nil, "function", "Test_PetById")
	petMock := petServiceMock{}
	petResponse := &PetResponse{
		ID:   1,
		Name: "Nash",
	}
	petMock.On("GetPetById", 1).Return(*petResponse, nil)
	petRouter := NewPetRouter(logger, &petMock)

	r := setupRouter()
	v1 := r.Group("/v1")
	petRouter.PetRegister(v1.Group("/pets"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/pets/1", nil)
	r.ServeHTTP(w, req)

	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)

	// unmarshal to Pet struct for asserts.
	actualPetResponse := &PetResponse{}
	json.Unmarshal(w.Body.Bytes(), actualPetResponse)
	assert.Equal(t, petResponse.ID, actualPetResponse.ID)
	assert.Equal(t, petResponse.Name, actualPetResponse.Name)
}

func Test_PetByName(t *testing.T) {
	logger := log.New().With(nil, "function", "Test_PetByName")
	petMock := petServiceMock{}

	petResponses := make([]PetResponse, 1)
	petRespopnse := &PetResponse{
		ID:   15,
		Name: "Charles",
	}
	petResponses[0] = *petRespopnse

	petMock.On("GetPetByName", "Charles").Return(petResponses, nil)
	petRouter := NewPetRouter(logger, &petMock)

	r := setupRouter()
	v1 := r.Group("/v1")
	petRouter.PetRegister(v1.Group("/pets"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/pets?name=Charles", nil)
	r.ServeHTTP(w, req)

	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)

	// unmarshal to Pet struct for asserts.
	actualPetResponses := make([]PetResponse, 1)
	json.Unmarshal(w.Body.Bytes(), &actualPetResponses)
	assert.Equal(t, petRespopnse.ID, actualPetResponses[0].ID)
	assert.Equal(t, petRespopnse.Name, actualPetResponses[0].Name)
}
