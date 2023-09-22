package main

import (
	"api/internal/api"
	"api/internal/sccclient"
)

func main() {
	c, err := sccclient.NewSCCWorkerClient("ws://localhost:8000/socketcluster/")
	if err != nil {
		panic(err)
	}

	if err = c.Connect(); err != nil {
		panic(err)
	}

	api.Start(c)
}
