package info

import (
	"github.com/qiangxue/go-restful-api/pkg/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getAppInfo(t *testing.T) {
	logger := log.New().With(nil, "function", "Test_getAppInfo")
	infoMock := MockServicer{}
	info := &Info{"Test App", "Info App", "1.0.0", "", ""}
	infoMock.On("getAppInfo").Return(*info, nil)

	infoService := NewInfoService(logger)
	result, _ := infoService.getAppInfo()
	infoMock.AssertExpectations(t)
	infoMock.AssertNumberOfCalls(t, "getAppInfo", 1)

	assert.Equal(t, info.AppName, result.AppName)
	assert.Equal(t, info.Description, result.Description)
	assert.Equal(t, info.Version, result.Version)
}
