PROJECT_NAME := bot
REPO_NAME := "gitlab.com/thundersnake/${PROJECT_NAME}"
PKG_LIST := $(shell go list ${REPO_NAME}/... | grep -v /vendor/)

.PHONY: all dep lint test race msan

all: test lint race msan

lint: ## test
	@go get -u github.com/golang/lint/golint
	@${GOPATH}/bin/golint -set_exit_status ${PKG_LIST}

test: dep ## Run unittests
	@go test -short ${PKG_LIST}

race: dep ## Run data race detector
	@go test -race -short ${PKG_LIST}

msan: dep ## Run memory sanitizer
	@go test -msan -short ${PKG_LIST}

dep:
	@go get -u github.com/golang/dep/cmd/dep
	@dep ensure
	mkdir -p ${CI_PROJECT_DIR}/artifacts/${GOOS}_${GOARCH}/
