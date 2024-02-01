package main

import "wire-petclinic-service/internal/api/info"

func ProvideInfoService() info.Service {
	return info.NewInfoService(logger)
}
