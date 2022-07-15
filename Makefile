SHELL = /bin/bash

.PHONY: run
run:
	air -c .air.toml

.PHONY: build
build:
	@mkdir -p tmp || true
	@go build -v -o tmp/ .