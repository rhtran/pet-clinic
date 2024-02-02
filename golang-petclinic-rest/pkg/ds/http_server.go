package ds

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rhtran/golang-petclinic-rest/app"
)

type HttpServer struct {
	r *gin.Engine
}

func NewHttpServer(r *gin.Engine) *HttpServer {
	return &HttpServer{
		r: r,
	}
}

func (httpSrv *HttpServer) HttpRouter() *http.Server {
	srv := &http.Server{
		Addr:    app.Config.Server.HttpPort,
		Handler: httpSrv.r,
	}

	return srv
}
