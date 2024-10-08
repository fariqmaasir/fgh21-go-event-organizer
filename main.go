package main

import (
	"github.com/fariqmaasir/fgh21-go-event-organizer/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(corsMiddleware())
	r.Static("/images", "./images")

	routers.RouterCombine(r)
	r.Run("0.0.0.0:8888")
}
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

//apakah koneksi di db pool / tidak
