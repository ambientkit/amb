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
