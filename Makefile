HELP_CMD="How to use it \n \
	make [target] \n \
    \n \
	List of available targets: \n \
	- drop - Drop all the tables from the database \n \
	- up - Recreate all the tables in the database and populate them with some mock data \n \
	- test - Run all the unit tests \n \
	- templ - Generate all the go files from the templates"



all:help

templ:
	@/home/$(USER)/go/bin/templ generate

drop:
	@echo "Drop cmd"

up:
	@echo "Up cmd"

test:
	@go test -v ./pkg/...

help:
	@echo $(HELP_CMD)


.PHONY: all templ help drop up test run