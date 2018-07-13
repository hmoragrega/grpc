package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hmoragrega/grpc/protobuf"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"golang.org/x/net/context"
)

// Greeter struct for proto comm
type Greeter struct{}

// Hello function for saying hello
func (g *Greeter) Hello(ctx context.Context, req *greeter.HelloRequest, resp *greeter.HelloResponse) error {
	resp.Greeting = "Hello " + req.Name
	fmt.Println("Responing with " + resp.Greeting)

	return nil
}

func main() {

	consulAddress := os.Getenv("CONSUL_ADDRESS")
	if consulAddress == "" {
		consulAddress = "localhost"
	}

	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("1.0.1"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*5),
		micro.Registry(registry.NewRegistry(registry.Addrs(consulAddress))),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	greeter.RegisterGreeterHandler(service.Server(), new(Greeter))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
