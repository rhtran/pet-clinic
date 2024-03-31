package ds

import (
	"github.com/rhtran/golang-petclinic-service/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	r *gin.Engine
}

func NewHttpServer(r *gin.Engine) *HttpServer {
	return &HttpServer{
		r: r,
	}
}

// HttpRouter is a method on the HttpServer struct that sets up and returns a new http.Server instance.
// The http.Server instance uses the gin.Engine instance from the HttpServer struct to handle HTTP requests.
func (httpSrv *HttpServer) HttpRouter() *http.Server {
	srv := &http.Server{
		Addr:    app.Config.Server.HttpPort,
		Handler: httpSrv.r,
	}

	return srv
}
