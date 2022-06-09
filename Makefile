.DEFAULT_GOAL=help

test: # Run test suite
	go test ./...

cov: # Calculate code coverage
	go test -coverprofile=coverage.out ./...

cov-html: cov # Calculate code coverage and generate html
	go tool cover -html=coverage.out

help: # Show this help
	@egrep -h '\s#\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?# "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
