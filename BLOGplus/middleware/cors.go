package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func ors() gin.HandlerFunc {
	return func(c *gin.Context) {
		cors.New(cors.Config{
			//AllowAllOrigins: true,
			AllowOrigins:  []string{"*"},
			AllowMethods:  []string{"*"},
			AllowHeaders:  []string{"Origin"},
			ExposeHeaders: []string{"Content-length", "Authorization"},
			//AllowCredentials: true,
			MaxAge: 12 * time.Hour,
		})
	}
}
