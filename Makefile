.PHONY: all

up:
	@docker-compose up -d

stop:
	@docker-compose stop

down:
	@docker-compose down

logs:
	@docker-compose logs -f

enter:
	@docker exec -it grpc_$(service)_1 /bin/sh

build: build-dep build-server build-client

build-dep:
	@docker run -v `pwd`:/go/src/github.com/hmoragrega/grpc golang:1.10.3-alpine /bin/sh -c 'cd /go/src/github.com/hmoragrega/grpc \
	&& apk add --no-cache git curl \
	&& curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh \
	&& dep ensure -v'

build-server:
	@docker build -t hmoragrega/grpc-server:latest -f server/Dockerfile .

build-client:
	@docker build -t hmoragrega/grpc-client:latest -f client/Dockerfile .

proto:
	@protoc --micro_out=. --go_out=. protobuf/greeter.proto

deploy:
	@kubectl apply -f kubernetes/role.yaml
	@kubectl apply -f kubernetes/server-deployment.yaml
	@kubectl apply -f kubernetes/server-service.yaml
	@kubectl apply -f kubernetes/client-deployment.yaml