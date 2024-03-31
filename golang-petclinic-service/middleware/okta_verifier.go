package middleware

import (
	"github.com/gin-gonic/gin"
)

type AuthenServicer interface {
	RedirectURL(c *gin.Context)
	Verify(bearerToken string) error
}

func Authenticate(authSrv AuthenServicer) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		err := authSrv.Verify(authHeader)
		if err != nil {
			authSrv.RedirectURL(c)
			return
		}

		c.Next()
	}
}
