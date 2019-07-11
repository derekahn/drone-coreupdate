# Project related variables
PROJECTNAME=$(shell basename "$(PWD)")
M = $(shell printf "\033[34;1m▶\033[0m")
DONE="\n  $(M)  done ✨"

.PHONY: help
help: Makefile
	@echo "\n Choose a command to run in "$(PROJECTNAME)":\n"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

## build: Creates a docker image of the app
.PHONY: build
build:
	@echo "  $(M)  🔨 Building the 🐳 image...\n"
	docker build -t=$(PROJECTNAME):dev .
	@echo $(DONE)

## clean: Removes the recently built docker image
.PHONY: clean
clean:
	@echo "  $(M)  🧹 last 🐳 image...\n"
	docker image rm $(PROJECTNAME):dev
	@echo $(DONE)

## install: Installs 🐹 dependencies
.PHONY: install
install:
	@echo "  $(M)  Checking if there is any missing dependencies...\n"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get $(get) ./...
	@echo $(DONE)

## run: Runs the current docker image
.PHONY: run
run:
	@echo "  $(M)  🏃 Running the 🐳 image...\n"
	docker run -it --rm --name $(PROJECTNAME) -t $(PROJECTNAME):dev
	@echo $(DONE)

## shell: To be executed after `make run` to give you a shell into the running container
.PHONY: shell
shell:
	@echo "  $(M)		📞 ...\n"
	docker exec -it $(PROJECTNAME) sh
	@echo $(DONE)
