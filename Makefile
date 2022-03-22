# This Makefile is an easy way to run common operations.
# Execute commands like this:
# * make

# Load the environment variables.
-include .env

.PHONY: default
default: run

################################################################################
# Setup app
################################################################################

.PHONY: run
run:
	go run cmd/amb/main.go

# Run go mod tidy.
.PHONY: tidy
tidy:
	go mod tidy -compat=1.17

# Update Go dependencies.
.PHONY: update
update:
	go get -u -f -d ./...
	go mod tidy -compat=1.17

# Pass in ARGS.
# https://stackoverflow.com/a/14061796
ifeq (update-ambient,$(firstword $(MAKECMDGOALS)))
  ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(ARGS):;@:)
endif

# Update Ambient dependency.
.PHONY: update-ambient
update-ambient:
	go get github.com/ambientkit/ambient@${ARGS}
	go mod tidy -compat=1.17

# Pass in ARGS.
# https://stackoverflow.com/a/14061796
ifeq (update-plugin,$(firstword $(MAKECMDGOALS)))
  ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(ARGS):;@:)
endif

# Update Ambient plugin dependency.
.PHONY: update-plugin
update-plugin:
	go get github.com/ambientkit/plugin@${ARGS}
	go mod tidy -compat=1.17