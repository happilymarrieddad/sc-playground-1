package sccclient_test

import (
	"encoding/json"
	"time"

	. "goclient/internal/sccclient"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("SC worker", func() {
	var (
		api1 SCCWorkerClient
		api2 SCCWorkerClient
		api3 SCCWorkerClient
	)

	BeforeEach(func() {
		var err error
		api1, err = NewSCCWorkerClient("ws://localhost:8000/socketcluster/")
		Expect(err).To(BeNil())
		Expect(api1).NotTo(BeNil())
		api2, err = NewSCCWorkerClient("ws://localhost:8001/socketcluster/")
		Expect(err).To(BeNil())
		Expect(api2).NotTo(BeNil())
		api3, err = NewSCCWorkerClient("ws://localhost:8002/socketcluster/")
		Expect(err).To(BeNil())
		Expect(api3).NotTo(BeNil())

		Expect(api1.Connect()).To(BeNil())
		Expect(api2.Connect()).To(BeNil())
		Expect(api3.Connect()).To(BeNil())
	})

	It("should successfully publish and handle data", func() {
		ch := "some-channel"
		ch2 := "some-channel-2"
		var count int

		channelHandler := func(name string) func(channelName string, data any) {
			return func(channelName string, data any) {
				bts, err := json.Marshal(data)
				Expect(err).To(BeNil())
				Expect(string(bts)).To(Equal(`{"message":"tester"}`))
				count++
			}
		}

		api2.Subscribe(ch, channelHandler("api2"))
		api2.Subscribe(ch2, channelHandler("api2"))

		api3.Subscribe(ch, channelHandler("api3"))
		api3.Subscribe(ch2, channelHandler("api3"))

		time.Sleep(time.Second * 1)
		api1.Publish(ch, struct {
			Message string `json:"message"`
		}{
			Message: "tester",
		})
		time.Sleep(time.Second * 2)
		Expect(count).To(BeNumerically("==", 2))
	})
})
