package info

import (
	"github.com/qiangxue/go-restful-api/pkg/log"
	"github.com/rhtran/golang-petclinic-rest/app"
)

type InfoService struct {
	logger log.Logger
}

func NewInfoService(logger log.Logger) *InfoService {
	return &InfoService{logger}
}

func (service *InfoService) GetAppInfo() (*Info, error) {
	// this block make GetAppInfo() not testable.
	info := &Info{
		AppName:     app.Config.AppInfo.Name,
		Description: app.Config.AppInfo.Description,
		Version:     app.Config.AppInfo.Version,
	}

	service.logger.Info(info)
	return info, nil
}
