.PHONY: build run

all: build

build:
	go build -o synop_cli cmd/synop/main.go

.PHONY: docs docs-deploy

docs:
	mkdocs serve

docs-deploy:
	mkdocs gh-deploy
