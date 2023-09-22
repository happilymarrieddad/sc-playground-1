package main

import (
	"goclient/internal/sccclient"
	"goclient/internal/service"
)

func main() {
	c, err := sccclient.NewSCCWorkerClient("ws://localhost:8000/socketcluster/")
	if err != nil {
		panic(err)
	}

	if err = c.Connect(); err != nil {
		panic(err)
	}

	service.Start(c)
}
