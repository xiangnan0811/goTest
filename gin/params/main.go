package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/welcome", welcome)
	r.POST("/form", form)
	r.POST("/post", getPost)
	r.Run()
}

func getPost(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.PostForm("name")
	message := c.DefaultPostForm("message", "Hello World")
	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"page":    page,
		"name":    name,
		"message": message,
	})
}

func form(c *gin.Context) {
	type Message struct {
		Name    string `form:"name" binding:"required"`
		Message string `form:"message" binding:"required"`
	}
	var msg Message
	if c.ShouldBind(&msg) == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": msg.Message,
			"name":    msg.Name,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "not posted",
		})
	}
}

func welcome(c *gin.Context) {
	first := c.DefaultQuery("first", "Guest")
	last := c.Query("last")
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome " + first + " " + last,
	})
}
