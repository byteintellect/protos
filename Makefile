PWD ?= $(shell pwd)
GO_CODE_PATH ?= "github.com/byteintellect/protos_go"
SOURCE_CODE_PATH ?= "${HOME}/go_proto_files"

.PHONY: compile_go
compile_go:
	. ./build.sh; build

.PHONY: push_go
push_go:
	. ./build.sh; push

.PHONY: build_and_push_go
build_and_push_go:
	. ./build.sh; build_and_push
