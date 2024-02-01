package info

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"wire-petclinic-service/internal/api/test"

	"github.com/gin-gonic/gin"

	"github.com/qiangxue/go-restful-api/pkg/log"

	"github.com/stretchr/testify/mock"
)

type infoServiceMock struct {
	mock.Mock
}

func (ism *infoServiceMock) GetAppInfo() (*Info, error) {
	args := ism.Called()
	intf := args.Get(0)
	val := intf.(Info)
	return &val, args.Error(1)
}

// config the gin engine for testing purpose
func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	return r
}

func TestInfoRouter(t *testing.T) {
	infoMock := infoServiceMock{}
	info := &Info{"Test App", "Info App", "1.0.0", "127.0.0.1", ""}
	infoMock.On("GetAppInfo").Return(*info, nil)

	logger := log.New().With(context.TODO(), "version", "test")
	r := setupRouter()
	home := r.Group("/")

	router := NewInfoRouter(logger, &infoMock, nil)
	router.InfoRegister(home.Group("/info"))

	tc1, _ := json.Marshal(info)

	tests := []test.APITestCase{
		{"get app info", "GET", "/info", "", nil, http.StatusOK, string(tc1)},
	}
	for _, tc := range tests {
		test.Endpoint(t, r, tc)
	}
}
