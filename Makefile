TOOL_DIR = .tool
BINPACK = $(TOOL_DIR)/binpack
TASK = $(TOOL_DIR)/task

.DEFAULT_GOAL := make-default

## Bootstrapping targets #################################

# note: we need to assume that binpack and task have not already been installed
$(BINPACK):
	@mkdir -p $(TOOL_DIR)
	@curl -sSfL https://raw.githubusercontent.com/khulnasoft/binpack/main/install.sh | sh -s -- -b $(TOOL_DIR)

# note: we need to assume that binpack and task have not already been installed
.PHONY: task
$(TASK) task: $(BINPACK)
	@$(BINPACK) install task -q

# this is a bootstrapping catch-all, where if the target doesn't exist, we'll ensure the tools are installed and then try again
%:
	@make $(TASK)
	@$(TASK) $@

## Shim targets #################################

.PHONY: make-default
make-default: $(TASK)
	@# run the default task in the taskfile
	@$(TASK)

help: $(TASK)
	@$(TASK) -l
