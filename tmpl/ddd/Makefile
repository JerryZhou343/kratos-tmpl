API_PROTO_FILES=$(shell cd proto && find . -name *.proto)

.PHONY: all
# generate all
all: build

.PHONY: grpc
# generate grpc code
grpc:
	 $(shell mkdir -p genproto)
	 cd proto  && protoc.exe --proto_path=. \
           --go_out=../genproto \
           --go-grpc_out=../genproto \
           $(API_PROTO_FILES)


.PHONY: wire
# generate wire
wire:
	cd cmd/server && wire


.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

