package server

import (
	"fmt"
	"log"
	"server-gin/config"

	"github.com/gin-gonic/gin"
)

func HttpListen() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	log.Fatal(r.Run(fmt.Sprintf(":%s", config.GetConfig().Ports.Http)))
}
