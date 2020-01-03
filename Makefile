.DEFAULT_GOAL := build

.PHONY: build
build:
	GOOS=linux go build -ldflags="-s -w"