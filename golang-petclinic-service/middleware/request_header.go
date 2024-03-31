package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

const (
	CORRELATION_HEADER = "X-Correlation-ID"
	REQUEST_ID_HEADER  = "X-Request-Id"
)

// SetRequestUUID will search for a correlation header and set a request-level
// correlation ID into the net.Context. If no header is found, a new UUID will
// be generated.
func SetRequestUUID() gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.Request.Header.Get(CORRELATION_HEADER)
		if u == "" {
			uuid, _ := uuid.NewV4()
			u = uuid.String()
		}
		c.Writer.Header().Set(CORRELATION_HEADER, u)

		uuid, _ := uuid.NewV4()
		c.Writer.Header().Set(REQUEST_ID_HEADER, uuid.String())
		c.Next()
	}
}
