FROM golang:1.22.4

WORKDIR src

RUN apt-get update && apt-get install -y protobuf-compiler

COPY go.mod .

RUN go env -w GOPROXY=https://goproxy.cn,direct

RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.4

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34

RUN go mod download