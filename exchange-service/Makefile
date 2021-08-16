PHONY: clean test build

clean:
	@rm -f exchange
	@echo "Cleaned!"

build:
	@go build -ldflags="-s -w" -o exchange ./main.go
	@echo "Binary generated."

run:
	@go run ./main.go

test:
	@gotestsum -f testname ./tests/...

lint:
	@staticcheck ./...

docker:
	@docker build -t exchange -f deployments/Dockerfile .
	@echo "Docker image generated."

prune:
	@docker image prune --filter label=stage=builder
	@echo "Removed build intermediate images."
