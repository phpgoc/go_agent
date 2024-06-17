FROM golang:1.22.4

WORKDIR src

COPY go.mod .

RUN go env -w GOPROXY=https://goproxy.cn,direct

RUN go mod download