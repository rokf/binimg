.DEFAULT_GOAL := help

.PHONY: build install help uninstall

help: ## Prints a descriptive target list
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-10s\033[0m %s\n", $$1, $$2}'

build: ## Builds the executable
	go build -o binimg binimg.go

install: ## Installs the executable to Go's binary installation location
	go install

uninstall: ## Removes the installed executable
	rm -i $(shell which binimg)
