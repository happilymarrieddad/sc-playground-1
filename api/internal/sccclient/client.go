package sccclient

import (
	"fmt"
	"sync"

	"github.com/happilymarrieddad/socketcluster-client-go/scclient"
)

type SCCWorkerClient interface {
	Connect() error
	Subscribe(channelName string, cb func(channelName string, data interface{}))
	Unsubscribe(channelName string)
	Publish(channelName string, data interface{})
	PublishAck(channelName string, data interface{}, cb func(channelName string, error interface{}, data interface{}))
}

func NewSCCWorkerClient(sscWorkerURL string) (SCCWorkerClient, error) {
	c := &sccWorkerClient{}

	c.cl = scclient.New(sscWorkerURL)

	return c, nil
}

type sccWorkerClient struct {
	cl scclient.Client
}

func (c *sccWorkerClient) Connect() (err error) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	c.cl.SetBasicListener(
		// onConnect
		func(client scclient.Client) {
			fmt.Println("api client connected to scc worker")
		},
		// onConnectError
		func(client scclient.Client, e error) {
			fmt.Printf("Error: %s\n", e.Error())
			err = e
			wg.Done()
		},
		// onDisconnect
		func(client scclient.Client, e error) {
			fmt.Printf("Error: %s\n", e.Error())
		})
	c.cl.SetAuthenticationListener(
		// onSetAuthentication
		func(client scclient.Client, token string) {
			fmt.Println("Auth token received :", token)
		},
		// onAuthentication
		func(client scclient.Client, isAuthenticated bool) {
			fmt.Println("Client authenticated :", isAuthenticated)
			wg.Done()
		})
	go c.cl.Connect()
	wg.Wait()

	return
}

func (c *sccWorkerClient) Subscribe(channelName string, cb func(channelName string, data interface{})) {
	c.cl.Subscribe(channelName)
	c.cl.OnChannel(channelName, cb)
}

func (c *sccWorkerClient) Unsubscribe(channelName string) {
	c.cl.Unsubscribe(channelName)
}

func (c *sccWorkerClient) Publish(channelName string, data interface{}) {
	c.cl.Publish(channelName, data)
}

func (c *sccWorkerClient) PublishAck(channelName string, data interface{}, cb func(channelName string, error interface{}, data interface{})) {
	c.cl.PublishAck(channelName, data, cb)
}
