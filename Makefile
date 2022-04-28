
git_tag=$(shell git describe --tags $(git rev-list --tags --max-count=1))
date=$(shell date)

.PHONY: pre-commit
pre-commit: lint test

.PHONY: test
test:
	go clean -testcache
	go test ./... -race -covermode=atomic -coverprofile=coverage.out

.PHONY: lint
lint:
	golangci-lint run ./...
