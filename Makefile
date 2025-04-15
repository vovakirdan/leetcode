.PHONY: install-hooks check update-readme

install-hooks:
	@bash scripts/install_hooks.sh

check:
	@gofmt -l . | tee /dev/stderr | (! read)
	@go vet ./...

update-readme:
	@go run scripts/generate_readme/generate_readme.go

test:
	@go test -v ./...

new-solution, new-problem, new:
	@bash scripts/new_solution.sh $(number) $(title) $(level)

fetch-tags:
	@go run scripts/fetch_tags/fetch_tags.go
