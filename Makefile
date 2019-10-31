init:
	dep ensure
	cp config.example.json config.json

serve:
	go run cmd/eventapi/*.go

test:
	 go test internal/tests/main_test.go

fetch:
	go run cmd/eventfetcher/main.go