package main

import (
	"fmt"
	"time"

	"github.com/hmoragrega/grpc/protobuf/greeter"
	grpc "github.com/micro/go-grpc"
	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	_ "github.com/micro/go-plugins/selector/static"
	"golang.org/x/net/context"
)

// The Greeter API.
type Greeter struct{}

func callEvery(d time.Duration, greeter greeter.GreeterService, f func(greeter.GreeterService)) {
	for range time.Tick(d) {
		f(greeter)
	}
}

func hello(greeterService greeter.GreeterService) {
	// Call the greeter
	rsp, err := greeterService.Hello(context.TODO(), &greeter.HelloRequest{
		Name: "Hilari",
	})

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%s\n", rsp.Greeting)
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
	greeter := greeter.NewGreeterService("greeter", service.Client())
	callEvery(3*time.Second, greeter, hello)
}
