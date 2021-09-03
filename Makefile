##########
# Building
##########

build-go:
	GOOS=linux GOARCH=arm go build -v -o dist/verpi

#########
# Linting
#########

lint-golangci:
	golangci-lint run
lint-gomod:
	go mod tidy
	git diff --exit-code go.mod
	git diff --exit-code go.sum
lint-goreleaser:
	goreleaser check

#########
# Testing
#########

test-go:
	go get -v -t -d ./...
	go test -v ./...

##########
# Grouping
##########

# Testing
local-test: test-go
# Linting
local-lint: lint-golangci lint-goreleaser lint-hadolint lint-gomod
# Build
local-build: build-docker-prod build-docker-dev build-docker-dev-lint
deploy: build-go
	scp dist/verpi $(PI):/home/pi/verpi
