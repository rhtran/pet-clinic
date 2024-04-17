package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockAuthenServicer struct {
	RedirectURLCalled bool
	VerifyCalled      bool
	VerifyError       error
}

func (m *MockAuthenServicer) RedirectURL(c *gin.Context) {
	m.RedirectURLCalled = true
}

func (m *MockAuthenServicer) Verify(bearerToken string) error {
	m.VerifyCalled = true
	return m.VerifyError
}

func TestAuthenticateWithValidToken(t *testing.T) {
	mockAuthSrv := &MockAuthenServicer{}
	router := gin.Default()
	router.Use(Authenticate(mockAuthSrv))
	router.GET("/", func(c *gin.Context) {})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer valid_token")

	router.ServeHTTP(w, req)

	assert.True(t, mockAuthSrv.VerifyCalled)
	assert.False(t, mockAuthSrv.RedirectURLCalled)
}

func TestAuthenticateWithInvalidToken(t *testing.T) {
	mockAuthSrv := &MockAuthenServicer{
		VerifyError: errors.New("invalid token"),
	}
	router := gin.Default()
	router.Use(Authenticate(mockAuthSrv))
	router.GET("/", func(c *gin.Context) {})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer invalid_token")

	router.ServeHTTP(w, req)

	assert.True(t, mockAuthSrv.VerifyCalled)
	assert.True(t, mockAuthSrv.RedirectURLCalled)
}

func TestAuthenticateWithNoToken(t *testing.T) {
	mockAuthSrv := &MockAuthenServicer{}
	router := gin.Default()
	router.Use(Authenticate(mockAuthSrv))
	router.GET("/", func(c *gin.Context) {})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	router.ServeHTTP(w, req)

	assert.False(t, mockAuthSrv.VerifyCalled)
	assert.True(t, mockAuthSrv.RedirectURLCalled)
}
