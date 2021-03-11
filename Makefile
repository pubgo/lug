Project=github.com/pubgo/golug
GOPath=$(shell go env GOPATH)
Version=$(shell git tag --sort=committerdate | tail -n 1)
GoROOT=$(shell go env GOROOT)
BuildTime=$(shell date "+%F %T")
CommitID=$(shell git rev-parse HEAD)
LDFLAGS=-ldflags " \
-X '${Project}/golug_version.GoROOT=${GoROOT}' \
-X '${Project}/golug_version.BuildTime=${BuildTime}' \
-X '${Project}/golug_version.GoPath=${GOPath}' \
-X '${Project}/golug_version.CommitID=${CommitID}' \
-X '${Project}/golug_version.Project=${Project}' \
-X '${Project}/golug_version.Version=${Version:-v0.0.1}' \
"

.PHONY: build
build:
	@go build ${LDFLAGS} -mod vendor -v -o main cmds/golug/main.go

build_hello_test:
	@go build ${LDFLAGS} -mod vendor -v -o main  example/hello/main.go

.PHONY: install
install:
	@cd cmds/golug && go install -v ${LDFLAGS} .

.PHONY: test
test:
	@go test -race -v ./... -cover

.PHONY: proto
proto: clear gen
	protoc -I. \
   -I/usr/local/include \
   -I${GOPATH}/src \
   -I${GOPATH}/src/github.com/googleapis/googleapis \
   -I${GOPATH}/src/github.com/gogo/protobuf \
   --go_out=plugins=grpc:. \
   --grpc-gateway_out=. \
   --grpc-gateway_opt=paths=source_relative \
   --grpc-gateway_opt=logtostderr=true \
   --golug_out=. \
	example/proto/hello/*

	protoc -I. \
   -I/usr/local/include \
   -I${GOPATH}/src \
   -I${GOPATH}/src/github.com/googleapis/googleapis \
   -I${GOPATH}/src/github.com/gogo/protobuf \
   --go_out=plugins=grpc:. \
   --grpc-gateway_out=. \
   --grpc-gateway_opt=paths=source_relative \
   --grpc-gateway_opt=logtostderr=true \
   --golug_out=. \
	example/proto/login/*

.PHONY: clear
clear:
	rm -rf example/proto/*.go
	rm -rf example/proto/**/*.go

.PHONY: gen
gen:
	cd cmds/protoc-gen-golug && go install .

.PHONY: example
example:
	go build ${LDFLAGS} -mod vendor -v -o main example/main.go

docker:
	docker build -t golug .