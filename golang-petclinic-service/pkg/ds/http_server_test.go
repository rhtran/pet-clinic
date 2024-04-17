package ds

import (
	"github.com/gin-gonic/gin"
	"github.com/rhtran/golang-petclinic-service/app"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHttpServerCreation(t *testing.T) {
	router := gin.Default()
	httpServer := NewHttpServer(router)

	assert.NotNil(t, httpServer)
	assert.Equal(t, router, httpServer.r)
}

func TestHttpRouter(t *testing.T) {
	router := gin.Default()
	httpServer := NewHttpServer(router)
	server := httpServer.HttpRouter()

	assert.NotNil(t, server)
	assert.Equal(t, app.Config.Server.HttpPort, server.Addr)
	assert.Equal(t, router, server.Handler)
}
