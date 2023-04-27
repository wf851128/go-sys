/*
 * @Description:
 */
//
package main

import (
	"go-sys/controller"
	"go-sys/model"

	"github.com/gin-gonic/gin"
)

func main() {
	// improt gin and create a new router
	model.InitDB()
	r := gin.Default()
	r1 := r.Group("/")
	{
		r1.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "index",
			})
		})
		// 用户api
		r1.POST("/login", controller.Login)
		r1.POST("/register", controller.Register)
		r1 := r.Group("user")
		{
			r1.POST("/:id", controller.PostUser)
			r1.PUT("/:id", controller.UpdateUser)
			r1.DELETE("/:id", controller.DeleteUser)
			r1.POST("", controller.PostUsers)
		}
	}
	r.Run("127.0.0.1:8077")
}
