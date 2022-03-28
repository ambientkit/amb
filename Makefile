# This Makefile is an easy way to run common operations.
# Execute commands like this:
# * make

# Load the environment variables.
-include .env

.PHONY: default
default: run

################################################################################
# Common
################################################################################

# Test the repo.
.PHONY: test
test:
	go test -race ./...

################################################################################
# Dependency management
################################################################################

# Update Ambient dependencies.
.PHONY: update
update: update-ambient update-plugin update-away tidy

# Update Ambient dependency. Requires the repo to be local and in the same folder.
.PHONY: update-ambient
update-ambient:
	go get -u github.com/ambientkit/ambient@$(shell cd ../ambient && git rev-parse HEAD)

# Update Ambient dependency. Requires the repo to be local and in the same folder.
.PHONY: update-plugin
update-plugin:
	go get -u github.com/ambientkit/plugin@$(shell cd ../plugin && git rev-parse HEAD)

# Update Ambient dependency. Requires the repo to be local and in the same folder.
.PHONY: update-away
update-away:
	go get -u github.com/ambientkit/away@$(shell cd ../away && git rev-parse HEAD)

# Update all Go dependencies.
.PHONY: update-all
update-all: update-all-go tidy

# Update all Go dependencies.
.PHONY: update-all-go
update-all-go:
	go get -u -f -d ./...

# Run go mod tidy.
.PHONY: tidy
tidy:
	go mod tidy -compat=1.17

################################################################################
# Setup app
################################################################################

.PHONY: run
run:
	go run cmd/amb/main.go
