PACKAGES_PATH = $(shell go list -f '{{.Dir}}' ./...)

.PHONY: tidy
tidy:
	@echo  "=> Executing go mod tidy"
	@go mod tidy