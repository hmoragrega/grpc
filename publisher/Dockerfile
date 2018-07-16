FROM golang:1.10.3-alpine AS build-env

COPY . /go/src/github.com/hmoragrega/grpc

WORKDIR /go/src/github.com/hmoragrega/grpc/publisher
RUN go build -o /bin/publisher .

# final stage
FROM alpine

COPY --from=build-env /bin/publisher /bin

CMD [ "publisher" ]