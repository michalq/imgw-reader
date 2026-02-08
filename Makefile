.PHONY: all clean

all: clean synop_cli

clean:
	@rm ./synop_cli || true

synop_cli:
	@go build -o synop_cli cmd/synop/main.go

.PHONY: docs docs-deploy
docs:
	mkdocs serve

docs-deploy:
	mkdocs gh-deploy
