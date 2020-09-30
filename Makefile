.PHONY: lint
lint:
				gofmt -s -w rpn

.PHONY: build
build:
				go build -v main.go

.PHONY: test
test:
				go test -cover -v -race -timeout 30s ./rpn

.DEFAULT_GOAL := build
