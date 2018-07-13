# grpc


## Installation fo Protocol Buffers (Mac OSX)

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

