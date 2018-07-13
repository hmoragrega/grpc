FROM golang:1.10.3-alpine AS build-env

COPY . /go/src/github.com/hmoragrega/grpc

RUN apk add --no-cache curl git
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR /go/src/github.com/hmoragrega/grpc
RUN dep ensure

WORKDIR /go/src/github.com/hmoragrega/grpc/client
RUN go build -o /bin/client .

# final stage
FROM alpine

COPY --from=build-env /bin/client /bin

CMD [ "client" ]