package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to my API",
	})
}

func ReturnGreetings(c *gin.Context) {
	name := c.Param("name")
	fmt.Println(name)
	c.JSON(200, gin.H{
		"message": "Hello, welcome " + name,
	})
}
