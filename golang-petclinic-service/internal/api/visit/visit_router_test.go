package visit

import (
	"encoding/json"
	"github.com/rhtran/golang-petclinic-service/internal/api/pet"
	"github.com/rhtran/golang-petclinic-service/pkg/infra/repository/test"
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

func Test_VisitById(t *testing.T) {
	logger := log.New().With(nil, "function", "Test_VisitById")
	visitMock := MockServicer{}

	visitResponse := &Response{
		ID:          1,
		Description: "spayed",
		PetResponse: pet.Response{
			ID:   1,
			Name: "Neo",
		},
	}

	visitMock.On("getVisitById", 1).Return(visitResponse, nil)
	visitRouter := NewVisitRouter(logger, &visitMock)

	r := setupRouter()
	v1 := r.Group("/v1")
	visitRouter.VisitRegister(v1.Group("/visits"))

	tc1, _ := json.Marshal(visitResponse)

	tests := []test.APITestCase{
		{"Get Visit By ID", "GET", "/v1/visits/1", "", nil, http.StatusOK, string(tc1)},
	}
	for _, tc := range tests {
		test.Endpoint(t, r, tc)
	}
}
