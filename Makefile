HELP_CMD="How to use it \nmake [target] \n\
\n\
List of available targets: \n\
- drop - Drop the database test (synapse_db) \n \
- up - Create the database test if not exists otherwise he uses it \n \
- test - Run all the unit tests \n \
- templ - Generate all the go files from the templates"

MYSQL_UP="CREATE DATABASE IF NOT EXISTS synapse_db; CREATE USER IF NOT EXISTS synapse@localhost IDENTIFIED BY 'SynapseT3stDb!'; GRANT ALL ON synapse_db.* TO 'synapse'@'localhost'; FLUSH PRIVILEGES;"

MYSQL_DROP="REVOKE ALL ON synapse_db.* FROM synapse@localhost; FLUSH PRIVILEGES; DROP USER IF EXISTS synapse@localhost; DROP DATABASE IF EXISTS synapse_db;"

all: help

templ:
	@/home/$(USER)/go/bin/templ generate

drop:
	@sudo mysql -e $(MYSQL_DROP)

up:
	@sudo mysql -e $(MYSQL_UP)

test:
	@go test -v ./pkg/...

help:
	@echo $(HELP_CMD)


.PHONY: all templ help drop up test run