package info

import (
	"github.com/qiangxue/go-restful-api/pkg/log"
	"net"
)

// IPServicer
type IPServicer interface {
	lookupIP(host string) ([]net.IP, error)
}

type IPService struct {
	logger log.Logger
}

func NewIPService(logger log.Logger) *IPService {
	return &IPService{logger}
}

func (service *IPService) lookupIP(host string) ([]net.IP, error) {
	return net.LookupIP(host)
}
