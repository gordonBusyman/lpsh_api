.PHONY: build test clean

# The binary to build (just the basename).
BIN := lpsh_api

# Where to build the binary.
OUTDIR := ./bin

# This version-strategy uses git tags to set the version string
VERSION := $(shell git describe --tags --always --dirty)

build:
	@echo "building ${VERSION}"
	@go build -o ${OUTDIR}/${BIN} -ldflags "-X main.Version=${VERSION}"

#test:
#	@go test -v ./...

clean:
	@echo "cleaning"
	@go clean
	@rm -rf ${OUTDIR}