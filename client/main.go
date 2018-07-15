package main

import (
	"fmt"
	"time"

	"github.com/hmoragrega/grpc/protobuf"
	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	"golang.org/x/net/context"
)

// The Greeter API.
type Greeter struct{}

func callEvery(d time.Duration, greeter greeter.GreeterService, f func(time.Time, greeter.GreeterService)) {
	for x := range time.Tick(d) {
		f(x, greeter)
	}
}

func hello(t time.Time, greeterService greeter.GreeterService) {
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
	service := micro.NewService(
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
