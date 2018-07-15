package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/hmoragrega/grpc/protobuf"
	micro "github.com/micro/go-micro"

	_ "github.com/micro/go-plugins/registry/kubernetes"
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

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("1.0.1"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	greeter.RegisterGreeterHandler(service.Server(), new(Greeter))

	go func() {
		http.HandleFunc("/health", health)
		livenessPort := os.Getenv("LIVENESS_PORT")
		if livenessPort == "" {
			livenessPort = "8080"
		}
		address := fmt.Sprintf("0.0.0.0:%s", livenessPort)
		fmt.Printf("Listening for liveness at %s\n", address)
		if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", livenessPort), nil); err != nil {
			log.Fatal(err)
		}
	}()

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
