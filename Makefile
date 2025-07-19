GOLANGCI_LINT := $(shell command -v golangci-lint 2>/dev/null)

.PHONY: lint test run

lint:
	@if [ -z "$(GOLANGCI_LINT)" ]; then \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi
	$(GOLANGCI_LINT) run

run:
	go run ./cmd/cats_api

test:
	go test ./...
