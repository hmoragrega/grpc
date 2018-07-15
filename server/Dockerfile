FROM golang:1.10.3-alpine AS build-env

COPY . /go/src/github.com/hmoragrega/grpc

WORKDIR /go/src/github.com/hmoragrega/grpc/server
RUN go build -o /bin/server .

# final stage
FROM alpine

COPY --from=build-env /bin/server /bin

CMD ["server"]