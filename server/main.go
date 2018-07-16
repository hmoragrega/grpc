package main

import (
	"fmt"
	"os"

	"github.com/hmoragrega/grpc/protobuf/greeter"

	grpc "github.com/micro/go-grpc"
	micro "github.com/micro/go-micro"

	_ "github.com/micro/go-plugins/broker/kafka"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	_ "github.com/micro/go-plugins/selector/static"
	"golang.org/x/net/context"
)

// Greeter struct for proto comm
type Greeter struct{}

// Hello function for saying hello
func (g *Greeter) Hello(ctx context.Context, req *greeter.HelloRequest, resp *greeter.HelloResponse) error {
	host, _ := os.Hostname()
	resp.Greeting = fmt.Sprintf("Hello %s from %s", req.Name, host)
	fmt.Println("Responing with " + resp.Greeting)

	return nil
}

// ProcessEvent consumes the events in the queue
func ProcessEvent(ctx context.Context, event *greeter.HelloRequest) error {
	fmt.Printf("Got event with name %s\n", event.Name)
	return nil
}

func main() {
	// Create a new service. Optionally include some options here.
	service := grpc.NewService(
		micro.Name("greeter"),
		micro.Version("1.0.1"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register greeter handler
	greeter.RegisterGreeterHandler(service.Server(), new(Greeter))

	// Register the subscriber
	micro.RegisterSubscriber("events", service.Server(), ProcessEvent)

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
