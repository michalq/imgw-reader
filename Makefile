.PHONY: build run
build:
	go build cmd/unzip/main.go
run:
	./main

.PHONY: docs docs-build
docs:
	mkdocs serve

docs-build:
	mkdocs build
