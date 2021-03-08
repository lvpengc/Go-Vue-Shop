
GOPATH:=$(shell go env GOPATH)


.PHONY: build
build:

	go build -o app_api-web *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t app_api-web:latest
