package api

import (
	"api/internal/sccclient"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
)

type TestMessage struct {
	Message string `json:"message"`
}

func Start(c sccclient.SCCWorkerClient) error {
	r := gin.Default()

	r.GET("/healtz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})

	c.Subscribe("test", func(channelName string, data interface{}) {
		spew.Dump(data)
	})

	r.Run()

	return nil
}
