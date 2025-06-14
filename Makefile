# Common makefile commands & variables between projects
include .make/common.mk

# Common Golang makefile commands & variables between projects
include .make/go.mk

## Not defined? Use default repo name which is the application
ifeq ($(REPO_NAME),)
	REPO_NAME="go-countries"
endif

## Not defined? Use default repo owner
ifeq ($(REPO_OWNER),)
	REPO_OWNER="mrz1836"
endif

.PHONY: all
all: ## Runs multiple commands
	@$(MAKE) test
