serve:
	go run cmd/eventapi/*.go

test:
	 go test internal/tests/main_test.go

fetch:
	go run cmd/eventfetcher/main.go