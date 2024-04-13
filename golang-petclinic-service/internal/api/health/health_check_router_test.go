package health

import (
	"context"
	"github.com/rhtran/golang-petclinic-service/internal/api/test"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/qiangxue/go-restful-api/pkg/log"
	"github.com/stretchr/testify/mock"
)

// healthServiceMock is a mock implementation of the HealthService interface.
// It's used for testing purposes.
type healthServiceMock struct {
	mock.Mock
}

// Check is a mock method that simulates the Check method of the HealthService interface.
// It returns a pointer to a Check object and an error.
// This is handwritten and used for testing purposes.
// The mockery library generates the mock implementation of the Check method.
func (hsm *healthServiceMock) check() (*Check, error) {
	args := hsm.Called()
	intf := args.Get(0)
	val := intf.(Check)
	return &val, args.Error(1)
}

// TestHealthRouter tests the HealthCheckRouter.
func TestHealthRouter(t *testing.T) {
	logger := log.New().With(context.TODO(), "version", "test")
	r := gin.Default()
	home := r.Group("/")

	healthMock := healthServiceMock{}
	healthUp := &Check{"UP"}
	healthMock.On("check").Return(*healthUp, nil)

	router := NewHealthCheckRouter(&healthMock, logger)
	router.HealthCheckRegister(home.Group("/health"))

	tests := []test.APITestCase{
		{"get app heartbeat", "GET", "/health", "", nil, http.StatusOK, `*UP*`},
	}
	for _, tc := range tests {
		test.Endpoint(t, r, tc)
	}
}

func TestHealthRouter_IsDown(t *testing.T) {
	logger := log.New().With(context.TODO(), "version", "test")
	r := gin.Default()
	home := r.Group("/")

	healthMock := healthServiceMock{}
	healthDown := &Check{"DOWN"}
	healthMock.On("check").Return(*healthDown, nil)

	router := NewHealthCheckRouter(&healthMock, logger)
	router.HealthCheckRegister(home.Group("/health"))

	tests := []test.APITestCase{
		{"get app heartbeat", "GET", "/health", "", nil, http.StatusOK, `*DOWN*`},
	}
	for _, tc := range tests {
		test.Endpoint(t, r, tc)
	}
}
