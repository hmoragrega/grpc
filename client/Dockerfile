FROM golang:1.10.3-alpine AS build-env

COPY . /go/src/github.com/hmoragrega/grpc

WORKDIR /go/src/github.com/hmoragrega/grpc/client
RUN go build -o /bin/client .

# final stage
FROM alpine

COPY --from=build-env /bin/client /bin

CMD [ "client" ]