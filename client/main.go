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

// The Greeter API.
type Greeter struct{}

// Hello is a Greeter API method.
func (g *Greeter) Hello(ctx context.Context, req *greeter.HelloRequest, rsp *greeter.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func callEvery(d time.Duration, greeter greeter.GreeterService, f func(time.Time, greeter.GreeterService)) {
	for x := range time.Tick(d) {
		f(x, greeter)
	}
}

func hello(t time.Time, greeterService greeter.GreeterService) {
	// Call the greeter
	rsp, err := greeterService.Hello(context.TODO(), &greeter.HelloRequest{
		Name: "Leander, calling at " + t.String(),
	})

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%s\n", rsp.Greeting)
	}
}

func main() {

	consulAddress := os.Getenv("CONSUL_ADDRESS")
	if consulAddress == "" {
		consulAddress = "localhost"
	}

	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),
		micro.Registry(registry.NewRegistry(registry.Addrs(consulAddress))),
	)

	// Init will parse the command line flags. Any flags set will
	// override the above settings.
	service.Init()

	// Create new greeter client and call hello
	greeter := greeter.NewGreeterService("greeter", service.Client())
	callEvery(3*time.Second, greeter, hello)
}
