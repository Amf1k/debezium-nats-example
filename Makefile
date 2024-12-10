export GOBIN ?= $(abspath bin)
export PATH := $(PATH):$(GOBIN)

NATS_CLI_CMD = $(GOBIN)/nats

$(NATS_CLI_CMD): go.mod
	@go install github.com/nats-io/natscli/nats

.PHONY: context
context: $(NATS_CLI_CMD) ## Setup localhost context
	@$(NATS_CLI_CMD) context add localhost --description "Localhost"

.PHONY: stream-ls
stream-ls: $(NATS_CLI_CMD)
	@$(NATS_CLI_CMD) stream ls

