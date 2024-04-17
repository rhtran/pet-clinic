package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetRequestUUIDWithCorrelationHeader(t *testing.T) {
	router := gin.Default()
	router.Use(SetRequestUUID())
	router.GET("/", func(c *gin.Context) {})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	correlationID := uuid.Must(uuid.NewV4()).String()
	req.Header.Set(CORRELATION_HEADER, correlationID)

	router.ServeHTTP(w, req)

	assert.Equal(t, correlationID, w.Header().Get(CORRELATION_HEADER))
	assert.NotEmpty(t, w.Header().Get(REQUEST_ID_HEADER))
}

func TestSetRequestUUIDWithoutCorrelationHeader(t *testing.T) {
	router := gin.Default()
	router.Use(SetRequestUUID())
	router.GET("/", func(c *gin.Context) {})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	router.ServeHTTP(w, req)

	assert.NotEmpty(t, w.Header().Get(CORRELATION_HEADER))
	assert.NotEmpty(t, w.Header().Get(REQUEST_ID_HEADER))
}
