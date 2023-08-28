package middleware

import (
	"GO_Practice/controllers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// This has to be checked,.....
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !controllers.Logged {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			fmt.Println("finished...")
			return
		}
		c.Next()
	}
}
