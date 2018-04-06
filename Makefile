#!/usr/bin/make
# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

PACKAGE = github.com/hawky-4s-/octoman
COMMIT_HASH = `git rev-parse --short HEAD 2>/dev/null`
BUILD_DATE = `date +%FT%T%z`
LDFLAGS = -ldflags "-X ${PACKAGE}/helpers.CommitHash=${COMMIT_HASH} -X ${PACKAGE}/helpers.BuildDate=${BUILD_DATE}"
NOGI_LDFLAGS = -ldflags "-X ${PACKAGE}/helpers.BuildDate=${BUILD_DATE}"

.PHONY: vendor docker check fmt lint test test-race vet test-cover-html help
.DEFAULT_GOAL := help

generate: ## Generate accessors
	go generate ../camunda-external-task-client-go/...

vendor: ## Install govendor and sync Octoman's vendored dependencies
	go get github.com/kardianos/govendor
	govendor sync ${PACKAGE}

octoman: vendor ## Build octoman binary
	go build ${LDFLAGS} ${PACKAGE}

octoman-race: vendor ## Build octoman binary with race detector enabled
	go build -race ${LDFLAGS} ${PACKAGE}

install: vendor ## Install octoman binary
	go install ${LDFLAGS} ${PACKAGE}

octoman-no-gitinfo: LDFLAGS = ${NOGI_LDFLAGS}
octoman-no-gitinfo: vendor octoman ## Build octoman without git info

docker: ## Build octoman Docker container
	docker build -t octoman .
	docker rm -f octoman-build || true
	docker run --name octoman-build octoman ls /go/bin
	docker cp octoman-build:/go/bin/octoman .
	docker rm octoman-build

check: test-race test386 fmt vet ## Run tests and linters

test386: ## Run tests in 32-bit mode
	GOARCH=386 govendor test +local

test: ## Run tests
	govendor test +local

test-race: ## Run tests with race detector
	govendor test -race +local

fmt: ## Run gofmt linter
	@for d in `govendor list -no-status +local | sed 's/github.com.hawky-4s-.octoman/./'` ; do \
		if [ "`gofmt -l $$d/*.go | tee /dev/stderr`" ]; then \
			echo "^ improperly formatted go files" && echo && exit 1; \
		fi \
	done

lint: ## Run golint linter
	@for d in `govendor list -no-status +local | sed 's/github.com.hawky-4s-.octoman/./'` ; do \
		if [ "`golint $$d | tee /dev/stderr`" ]; then \
			echo "^ golint errors!" && echo && exit 1; \
		fi \
	done

vet: ## Run go vet linter
	@if [ "`govendor vet +local | tee /dev/stderr`" ]; then \
		echo "^ go vet errors!" && echo && exit 1; \
	fi

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
