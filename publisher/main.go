package main

import (
	"fmt"
	"time"

	"github.com/hmoragrega/grpc/protobuf/greeter"
	grpc "github.com/micro/go-grpc"
	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/broker/kafka"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	_ "github.com/micro/go-plugins/selector/static"
	"golang.org/x/net/context"
)

// The Greeter API.
type Greeter struct{}

func callEvery(d time.Duration, publisher micro.Publisher, f func(micro.Publisher)) {
	for range time.Tick(d) {
		f(publisher)
	}
}

func hello(publisher micro.Publisher) {
	err := publisher.Publish(context.TODO(), &greeter.HelloRequest{
		Name: "Hilari",
	})

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Event published")
	}
}

func main() {
	// Create a new service. Optionally include some options here.
	service := grpc.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),
	)

	// Init will parse the command line flags. Any flags set will
	// override the above settings.
	service.Init()

	// Create new greeter client and call hello
	publisher := micro.NewPublisher("events", service.Client())
	callEvery(3*time.Second, publisher, hello)
}
