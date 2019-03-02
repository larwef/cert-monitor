VERSION=v0.0.1
SOURCE=cmd
TARGET=target
LAMBDA_TARGET=$(TARGET)/lambda/certsearch-$(VERSION)-lambda-deployment.zip
WEBAPP_TARGET=$(TARGET)/webapp
REPO=repo.wefald.no/certSearch
PROFILE=larwef

all: build-webapp serve

build: build-lambda build-webapp

build-lambda:
	GOOS=linux go build -o $(TARGET)/lambda/main $(SOURCE)/lambda/main.go
	zip -j $(LAMBDA_TARGET) $(TARGET)/lambda/main

build-webapp:
	GOOS=js GOARCH=wasm go build -o $(WEBAPP_TARGET)/main.wasm $(SOURCE)/webapp/*.go
	cp -r web/ $(WEBAPP_TARGET) 

serve:
	go run cmd/server/server.go -dir target/webapp/

integration:
	go test test/integration/handler/handler_test.go -v -tags=integration

client:
	go run cmd/client/main.go

upload:
	aws s3 cp $(LAMBDA_TARGET) s3://$(REPO)/lambda/certsearch-$(VERSION)-lambda-deployment.zip --profile $(PROFILE)
	aws s3 cp $(WEBAPP_TARGET) s3://$(REPO)/webapp/ --recursive --profile $(PROFILE)