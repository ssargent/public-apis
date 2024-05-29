
# the version of protoc used to build protobuf
PROTOC_VERSION = 3.13.0

.PHONY: proto-lint
lint: 
	@buf lint

generate: proto-lint
	@buf generate