VERSION=v0.0.1
SOURCE=cmd
TARGET=target

all: build-webapp serve

build: build-lambda build-webapp

build-lambda:
	GOOS=linux go build -o $(TARGET)/lambda/main $(SOURCE)/lambda/main.go
	zip -j $(TARGET)/lambda/certsearch-$(VERSION)-lambda-deployment.zip $(TARGET)/lambda/main

build-webapp:
	GOOS=js GOARCH=wasm go build -o $(TARGET)/webapp/main.wasm $(SOURCE)/webapp/*.go
	cp -r web/ $(TARGET)/webapp 

serve:
	go run cmd/server/server.go -dir target/webapp/

integration:
	go test test/integration/handler/handler_test.go -v -tags=integration

client:
	go run cmd/client/main.go