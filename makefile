PACKAGES_PATH = $(shell go list -f '{{.Dir}}' ./...)

.PHONY: tidy
tidy:
	@echo  "=> Executing go mod tidy"
	@go mod tidy

.PHONY: fmt
fmt:
	@echo  "=> Executing go fmt"
	@go fmt ./...