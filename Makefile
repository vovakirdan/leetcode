.PHONY: install-hooks check update-readme

install-hooks:
	@bash scripts/install_hooks.sh

check:
	@gofmt -l . | tee /dev/stderr | (! read)
	@go vet ./...

update-readme:
	@go run scripts/generate_readme.go

test:
	@go test -v ./...

new-solution:
	@bash scripts/new_solution.sh $(number) $(title) $(level)