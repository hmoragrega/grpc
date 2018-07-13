.PHONY: all

PROTOSERVICE = protobuf/greeter.proto
PROTOC       = protoc
PROTOCOBJ    = --go_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:.

up:
	@docker-compose up -d

down:
	@docker-compose down

build:
	@docker-compose build

logs:
	@docker-compose logs -f

enter:
	@docker exec -it grpc_$(service)_1 /bin/sh

proto:
	@protoc --micro_out=. --go_out=. protobuf/greeter.proto