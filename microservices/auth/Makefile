# Definitions
GOPATH					?= $(shell go env GOPATH)
PATH 					:= $(PATH):$(GOPATH)/bin
GOROOT 					:= $(GOROOT)
ROOT                    = $(PWD)
GOLANG_DOCKER_IMAGE     = golang:alpine3.15
GOLANG_DOCKER_CONTAINER = empty
DOCKER_IMAGE_NAME		= auth-app-image
DOCKER_IMAGE_VERSION	= 1.0.0
GO_APP_NAME				= auth-app
GO_APP_PROD_FILE		= "cmd/prod/main.go"

clean-docker:	##deletes container if exists
	docker rm -f ${GOLANG_DOCKER_CONTAINER} || true

docker-run:	##run
	docker container run -it --rm -p 9091:9091 auth-app-image:1.0.0

go-build: go-proto	##build executable go file
	go version
	go mod download
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "exec/${GO_APP_NAME}" ./"${GO_APP_PROD_FILE}"

build-docker-image: go-build	##build docker image
	@docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_VERSION) .

go-proto: 	##generate proto files and copy them to go-app
	make -C ../../protolib run
	cp -r ../../protolib/gen/proto/go/* ./proto

help:	##Show this help.
	@grep -e "^[a-zA-Z_-]*:" Makefile|awk -F'##' '{gsub(/[ \t]+$$/, "", $$1);printf "%-30s\t%s\n", $$1, $$2}'


