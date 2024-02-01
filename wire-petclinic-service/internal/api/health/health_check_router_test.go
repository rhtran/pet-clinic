package health

import (
	"context"
	"log"
	"net/http"
	"testing"
	"wire-petclinic-servic/internal/api/test"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type healthServiceMock struct {
	mock.Mock
}

func (hsm *healthServiceMock) Check() (*HealthCheck, error) {
	args := hsm.Called()
	intf := args.Get(0)
	val := intf.(HealthCheck)
	return &val, args.Error(1)
}

func TestHealthRouter(t *testing.T) {
	healthMock := healthServiceMock{}
	health := &HealthCheck{"UP"}
	healthMock.On("Check").Return(*health, nil)

	logger := log.New().With(context.TODO(), "version", "test")
	r := gin.Default()
	home := r.Group("/")

	router := NewHealthCheckRouter(&healthMock, logger)
	router.HealthCheckRegister(home.Group("/health"))

	tests := []test.APITestCase{
		{"get app health", "GET", "/health", "", nil, http.StatusOK, `*UP*`},
	}
	for _, tc := range tests {
		test.Endpoint(t, r, tc)
	}
}
