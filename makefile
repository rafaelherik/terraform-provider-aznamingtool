# Define variables
TERRAFORM_CMD=terraform
ENV_SCRIPT=./examples/env.sh

# Default rule to display help
.PHONY: help
help:
	@echo "Usage:"
	@echo "  make run-test DIR=<<DIRECTORY WITH TERRAFORM FILES>>"	

# Rule to run tests
.PHONY: run-test
run-testr: check-dir
	@echo "Running tests in directory: $(DIR)"
	@source $(ENV_SCRIPT) && \
	cd $(DIR) && \
	$(TERRAFORM_CMD) init && \
	$(TERRAFORM_CMD) fmt -check && \
	$(TERRAFORM_CMD) plan && \
	$(TERRAFORM_CMD) apply -auto-approve && \
	$(TERRAFORM_CMD) destroy -auto-approve

# Rule to check if DIR variable is set
.PHONY: check-dir
check-dir:
ifndef DIR
	$(error DIR is undefined. Usage: make run-test DIR=./examples/test1)
endif
