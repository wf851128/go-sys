//
package main

import (
	"go-sys/model"

	"github.com/gin-gonic/gin"
)

func main() {
	// improt gin and create a new router
	model.InitDB()
	r := gin.Default()
	r1 := r.Group("")
	{
		r1.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "hello world",
			})
		})
	}

	r.Run("127.0.0.1:8077")
}
