package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xiangnan0811/goTest/proto"
)

func main() {
	r := gin.Default()
	r.GET("/proto", returnProto)
	r.Run(":8080")
}

func returnProto(c *gin.Context) {
	c.ProtoBuf(http.StatusOK, &proto.Teacher{
		Name:   "张三",
		Course: []string{"语文", "数学", "英语"},
	})
}