.PHONY: build run
build:
	go build cmd/unzip/main.go
run:
	./main

.PHONY: docs docs-deploy
docs:
	mkdocs serve

docs-deploy:
	mkdocs gh-deploy
