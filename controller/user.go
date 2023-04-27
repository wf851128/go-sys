package controller

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Login",
	})
}

func Register(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Register",
	})
}

func PostUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": id,
	})
}

func UpdateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "UpdateUser",
	})
}

func DeleteUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "DeleteUser",
	})
}

func PostUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "PostUsers",
	})
}
