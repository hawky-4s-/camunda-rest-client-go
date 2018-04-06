#!/usr/bin/make
# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

PACKAGE = github.com/hawky-4s-/camunda-rest-client-go
VERSION_GOMETALINDER := 2.0.5

.PHONY: vendor check fmt lint test test-race vet test-cover-html help integration-test performance-test
.DEFAULT_GOAL := help

generate: ## Generate accessors
	go generate ../camunda-rest-client-go/...

vendor: ## Install dep and get dependencies
	dep ensure

check: test-race fmt vet ## Run tests and linters

test386: ## Run tests in 32-bit mode
	GOARCH=386 govendor test +local

test: ## Run tests
	go test -race `go list -tags=unit ./... | grep -v examples`

test-all: test-unit test-integration test-performance ## Run all tests

test-unit: ## Run unit tests
	go test -tags=unit `go list -tags=unit ./... | grep -v examples`

test-integration: ## Run integration tests
	# start docker before
	go test -tags=integration `go list -tags=integration ./... | grep -v examples`
	# end docker

test-performance: ## Run performance tests
	go test -tags=performance `go list -tags=performance ./... | grep -v examples`

test-race: ## Run tests with race detector
	go test -race +local

fmt: ## Run gofmt linter
	@for d in `govendor list -no-status +local | sed 's/github.com.hawky-4s-.octoman/./'` ; do \
		if [ "`gofmt -l $$d/*.go | tee /dev/stderr`" ]; then \
			echo "^ improperly formatted go files" && echo && exit 1; \
		fi \
	done

lint: ## Run gometalinter
	curl -Lo https://github.com/alecthomas/gometalinter/releases/download/v$(VERSION_GOMETALINTER)/gometalinter-$(VERSION_GOMETALINTER)-$(shell sed -e 's/\(.*\)/\L\1/')-amd64.tar.gz

vet: ## Run go vet linter
	go vet ./...

test-cover-html: PACKAGES = $(shell govendor list -no-status +local | sed 's/github.com.hawky-4s-.octoman/./')
test-cover-html: ## Generate test coverage report
	echo "mode: count" > coverage-all.out
	$(foreach pkg,$(PACKAGES),\
		govendor test -coverprofile=coverage.out -covermode=count $(pkg);\
		tail -n +2 coverage.out >> coverage-all.out;)
	go tool cover -html=coverage-all.out

check-vendor: ## Verify that vendored packages match git HEAD
	@git diff-index --quiet HEAD vendor/ || (echo "check-vendor target failed: vendored packages out of sync" && echo && git diff vendor/ && exit 1)

help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
