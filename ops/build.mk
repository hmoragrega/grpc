# Build
build: build-deps build-server build-client

build-deps:
	@docker run -v `pwd`:/go/src/github.com/hmoragrega/grpc golang:1.10.3-alpine /bin/sh -c 'cd /go/src/github.com/hmoragrega/grpc \
	&& apk add --no-cache git curl \
	&& curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh \
	&& dep ensure -v'

build-server:
	@docker build -t hmoragrega/grpc-server:latest -f server/Dockerfile .

build-client:
	@docker build -t hmoragrega/grpc-client:latest -f client/Dockerfile .

build-proto:
	@protoc --micro_out=. --go_out=. protobuf/greeter.proto