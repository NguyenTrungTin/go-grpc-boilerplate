FROM golang:1.14.4-alpine

RUN apk update && apk upgrade && \
	apk add --no-cache bash git openssh

WORKDIR /app

RUN go get github.com/golang/protobuf/protoc-gen-go \
	github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
	github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

RUN go get -u github.com/mitranim/gow

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 9901 9902

CMD ["./main", "server"]
