package info

import (
	"context"
	"encoding/json"
	"github.com/rhtran/golang-petclinic-service/pkg/infra/repository/test"
	"net"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/qiangxue/go-restful-api/pkg/log"

	"github.com/stretchr/testify/mock"
)

type infoServiceMock struct {
	mock.Mock
}

func (ism *infoServiceMock) getAppInfo() (*Info, error) {
	args := ism.Called()
	intf := args.Get(0)
	val := intf.(Info)
	return &val, args.Error(1)
}

type ipServiceMock struct {
	mock.Mock
}

func (ips *ipServiceMock) lookupIP(host string) ([]net.IP, error) {
	args := ips.Called(host)
	intf := args.Get(0)
	val := intf.([]net.IP)
	return val, args.Error(1)
}

// config the gin engine for testing purpose
func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	return r
}

func TestInfoRouter(t *testing.T) {
	infoMock := infoServiceMock{}
	info := &Info{"Test App", "Info App", "1.0.0", "", ""}
	expectInfo := &Info{"Test App", "Info App", "1.0.0", "127.0.0.1", ""}
	infoMock.On("getAppInfo").Return(*info, nil)

	ipServiceMock := ipServiceMock{}
	ipServiceMock.On("lookupIP", "localhost").Return([]net.IP{net.ParseIP("127.0.0.1")}, nil)

	logger := log.New().With(context.TODO(), "version", "test")
	r := setupRouter()
	home := r.Group("/")

	router := NewInfoRouter(logger, &infoMock, &ipServiceMock)
	router.InfoRegister(home.Group("/info"))

	tc1, _ := json.Marshal(expectInfo)

	tests := []test.APITestCase{
		{"get app info", "GET", "/info", "", nil, http.StatusOK, string(tc1)},
	}
	for _, tc := range tests {
		test.Endpoint(t, r, tc)
	}
}
