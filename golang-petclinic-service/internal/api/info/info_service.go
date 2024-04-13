package info

import (
	"github.com/qiangxue/go-restful-api/pkg/log"
	"github.com/rhtran/golang-petclinic-service/app"
)

type Servicer interface {
	getAppInfo() (*Info, error)
}

type Service struct {
	logger log.Logger
}

func NewInfoService(logger log.Logger) *Service {
	return &Service{logger}
}

func (service *Service) getAppInfo() (*Info, error) {
	// this block make GetAppInfo() not testable.
	info := &Info{
		AppName:     app.Config.AppInfo.Name,
		Description: app.Config.AppInfo.Description,
		Version:     app.Config.AppInfo.Version,
	}

	service.logger.Info(info)
	return info, nil
}
