# Microservices Orchestration Demo
- Service discovery with either consul or kubernetes services
- Load balancing between nodes
- Pub/Sub Events with Kafka
- Circuit breaker [TODO]

## Build
Build docker images
```
make build
```

## Run
### Option 1: Docker + Consul
You can run the demo locally in docker using consul as service discovery (go-micro default registry)
```
make up
```
To check the output
```
make logs
```
To stop it type
```
make down
```
### Option 2: Kubernetes in minikube
You can use a local kubernetes cluster to deploy the demo
Install minikube an activate set up the terminal to use the minikube docker
```
eval $(minikube docker-env)
```
Rebuild the images so they are available inside the minikube docker
```
make build
```
Deploy Kafka, it will take some seconds
```
make deploy-kafka
```
Once kafka is ready then deploy the demo
```
make deploy
```
To remove all services type:
```
make destroy
make destroy-kafka
```

## Appendix: Installation fo Protocol Buffers (Mac OSX)

### Install the proto compiler

### With brew

**NOTE**: The next steps will install the compiler from the source, alternatively you can use brew
```
brew install protobuf
```

### From source

#### Prerequisites
```
brew install autoconf && brew install automake
```

1. Download the appropriate release here: https://github.com/google/protobuf/releases

2. Configure and build the compiler
```
./autogen.sh
./configure
make
make check
```

3. Install it
```
sudo make install
which protoc
protoc --version
```

## Install the Go plugin for the proto compiler
```
go get -u github.com/golang/protobuf/protoc-gen-go
```

Make sure `$GOPATH/bin` is available on the `$PATH`

## Compile the proto spec to Go classes
```
protoc --go_out=. protobuf/greeter.proto
```

