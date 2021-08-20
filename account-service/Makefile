run:
	@go run ./main.go

test:
	@gotestsum -f testname ./tests/...

int-test:
	@go test -v -tags integration ./tests/...

coverage:
	@go test -v -cover -coverprofile=r.out -coverpkg ./internal/... ./tests/...
	@go tool cover -html=r.out -o report.hmlt
	@rm -f r.out