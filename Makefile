# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## audit: run quality control checks
.PHONY: audit
audit:
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	go test -race -buildvcs -vet=off ./...


# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

## release: build the release version of the website
.PHONY: release
release:
	@go run ./cmd/gen public
	@cp -r assets/static public/static/

## build: build the development version of the website
.PHONY: build
build:
	@go run ./cmd/gen tmp
	@cp -r assets/static tmp/static/
	
## run: build and serve the website
.PHONY: run
run: build
	@go run ./cmd/serve tmp

## run/live: build and serve the website with reloading on file changes
.PHONY: run/live
run/live:
	@go run github.com/cosmtrek/air@v1.43.0

## clean: Clean generated files.
.PHONY: clean
clean:
	rm -rf public
	rm -rf tmp
