package mid_auth

import "github.com/gin-gonic/gin"

func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"user-rsud": "password123",
	})
}
