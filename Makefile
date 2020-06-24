#!make
include .env
# export $(shell sed 's/=.*//' .env)

install:
	go mod download
	go get github.com/mitranim/gow \
		github.com/golang/protobuf/protoc-gen-go \
    	github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
    	github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

gen:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb --grpc-gateway_out=:pb --swagger_out=logtostderr=true:docs/api

clean:
	rm pb/*.go
	rm docs/api/*.swagger.json

server:
	gow run main.go server --port ${GRPC_PORT} --http ${HTTP_PORT}

server-tls:
	go run main.go server --port ${GRPC_PORT} --http ${HTTP_PORT} --tls

client:
	gow run main.go client --address ${GRPC_PORT}

client-tls:
	go run main.go client --address ${GRPC_PORT} --tls

test:
	go test -cover -race ./...

grpc-cert:
	chmod +x hack/generate-grpc-certs.sh && hack/generate-grpc-certs.sh

.PHONY: gen clean server client test cert

