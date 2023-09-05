package main

import(
	"github.com/gin-gonic/gin"
)

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	g := gin.Default()
	g.GET("/ping", pong)
	g.Run(":8085")
}
