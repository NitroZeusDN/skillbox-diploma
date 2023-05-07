run_processor:
	go run ./cmd/processor/main.go

run_simulator:
	go run ./cmd/simulator/main.go

lint:
	golangci-lint run -v --timeout 15m --fix cmd/... internal/...